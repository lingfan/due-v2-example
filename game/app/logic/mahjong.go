package logic

import (
	"context"
	"due-v2-example/game/app/entity"
	"due-v2-example/shared/code"
	"due-v2-example/shared/middleware"
	"due-v2-example/shared/pb/common"
	pb "due-v2-example/shared/pb/mahjong"
	"due-v2-example/shared/route"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/config"
	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/session"
)

type Mahjong struct {
	proxy     *node.Proxy
	ctx       context.Context
	roomMgr   *entity.RoomMgr
	playerMgr *entity.PlayerMgr
}

func NewMahjong(proxy *node.Proxy) *Mahjong {
	opts := make([]*entity.Options, 0)
	err := config.Match("mahjong.rooms").Scan(&opts)
	log.Infof("NewMahjong config.Get %+v ", opts)
	if err != nil {
		log.Fatalf("load mahjong rooms config failed: %v", err)
	}

	return &Mahjong{
		proxy:     proxy,
		ctx:       context.Background(),
		roomMgr:   entity.NewRoomMgr(opts),
		playerMgr: entity.NewPlayerMgr(),
	}
}

func (l *Mahjong) Init() {
	// 注册路由
	l.proxy.Router().Group(func(group *node.RouterGroup) {
		// 注册中间件
		group.Middleware(middleware.Auth)
		// 快速开始
		group.AddRouteHandler(route.QuickStart, false, l.quickStart)
		// 坐下
		group.AddRouteHandler(route.SitDown, false, l.sitDown)
		// 起立
		group.AddRouteHandler(route.StandUp, true, l.standUp)
		// 开始准备
		group.AddRouteHandler(route.Ready, true, l.ready)
		// 取消准备
		group.AddRouteHandler(route.Unready, true, l.unready)
	})

	// 断开连接
	l.proxy.Events().AddEventHandler(cluster.Disconnect, l.disconnect)
	// 重新连接
	l.proxy.Events().AddEventHandler(cluster.Reconnect, l.reconnect)
}

// 断开连接
func (l *Mahjong) disconnect(event *node.Event) {
	player, err := l.playerMgr.GetPlayer(event.UID)
	if err != nil {
		return
	}

	seat := player.Seat()
	if seat == nil {
		return
	}

	seat.Offline()

	l.syncSeatStateChange(seat, pb.SeatState_Offline)
}

// 重新连接
func (l *Mahjong) reconnect(event *node.Event) {
	player, err := l.playerMgr.GetPlayer(event.UID)
	if err != nil {
		return
	}

	seat := player.Seat()
	if seat == nil {
		return
	}

	seat.Online()

	l.syncSeatStateChange(seat, pb.SeatState_Online)

	l.syncGameInfo(player)
}

// 快速开始
func (l *Mahjong) quickStart(ctx *node.Context) {
	res := &pb.QuickStartRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("quick start response failed, err: %v", err)
		}
	}()

	player, err := l.playerMgr.LoadPlayer(ctx.Request.UID)
	if err != nil {
		log.Errorf("load user info failed: uid: %d err: %v", ctx.Request.UID, err)
		res.Code = common.Code_Failed
		return
	}

	log.Infof("playerMgr.LoadPlayer %v %+v", ctx.Request.UID, player)

	seat := player.Seat()
	log.Infof("player.Seat %+v", seat)
	if seat != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	defer func() {
		if err != nil {
			l.playerMgr.UnloadPlayer(ctx.Request.UID)
		}
	}()

	if err = l.roomMgr.QuickMatch(player); err != nil {
		log.Errorf("quick match failed: uid: %d err: %v", ctx.Request.UID, err)
		res.Code = common.Code_Failed
		return
	}

	if err = ctx.BindNode(); err != nil {
		log.Errorf("bind node failed: uid: %d err: %v", ctx.Request.UID, err)
		res.Code = common.Code_Failed
		return
	}

	l.syncTakeSeatInfo(player.Seat())

	res.Code = common.Code_OK
	res.GameInfo = l.makeGameInfo(player)
}

// 坐下
func (l *Mahjong) sitDown(ctx *node.Context) {
	req := &pb.SitDownReq{}
	res := &pb.SitDownRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("sit down response failed, err: %v", err)
		}
	}()

	player, err := l.playerMgr.LoadPlayer(ctx.Request.UID)
	if err != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	seat := player.Seat()
	if seat != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	if err = ctx.Request.Parse(req); err != nil {
		log.Errorf("invalid sit down message, err: %v", err)
		res.Code = common.Code_Abnormal
		return
	}

	seat, err = l.roomMgr.GetSeat(int(req.RoomID), int(req.TableID), int(req.SeatID))
	if err != nil {
		res.Code = common.Code_IllegalParams
		return
	}

	defer func() {
		if err != nil {
			l.playerMgr.UnloadPlayer(ctx.Request.UID)
		}
	}()

	err = seat.AddPlayer(player)
	if err != nil {
		switch errors.Code(err) {
		case code.SeatAlreadyTaken:
			res.Code = common.Code_SeatAlreadyTaken
		case code.PlayerAlreadySeated:
			res.Code = common.Code_IllegalOperation
		default:
			res.Code = common.Code_Failed
		}
		log.Errorf("sit down failed: uid: %d roomID: %d tableID: %d seatID: %d err: %v", ctx.Request.UID, req.RoomID, req.TableID, req.SeatID, err)
		return
	}

	if err = ctx.BindNode(); err != nil {
		log.Errorf("bind node failed: uid: %d err: %v", ctx.Request.UID, err)
		res.Code = common.Code_Failed
		return
	}

	l.syncTakeSeatInfo(seat)

	res.Code = common.Code_OK
	res.GameInfo = l.makeGameInfo(player)
}

// 站起
func (l *Mahjong) standUp(ctx *node.Context) {
	res := &pb.StandUpRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("start ready response failed, err: %v", err)
		}
	}()

	if ctx.Request.UID <= 0 {
		res.Code = common.Code_NotLogin
		return
	}

	player, err := l.playerMgr.GetPlayer(ctx.Request.UID)
	if err != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	seat := player.Seat()
	if seat == nil {
		res.Code = common.Code_OK
		return
	}

	err = seat.RemPlayer()
	if err != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	err = ctx.UnbindNode()
	if err != nil {
		log.Errorf("unbind node failed: uid: %d err: %v", ctx.Request.UID, err)
	}

	l.syncSeatStateChange(seat, pb.SeatState_StandUp)

	res.Code = common.Code_OK
}

// 开始准备
func (l *Mahjong) ready(ctx *node.Context) {
	res := &pb.ReadyRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("start ready response failed, err: %v", err)
		}
	}()

	player, err := l.playerMgr.GetPlayer(ctx.Request.UID)
	if err != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	seat := player.Seat()
	if seat == nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	if seat.IsReady() {
		res.Code = common.Code_OK
		return
	}

	seat.Ready()

	l.syncSeatStateChange(seat, pb.SeatState_Ready)

	res.Code = common.Code_OK
}

// 取消准备
func (l *Mahjong) unready(ctx *node.Context) {
	res := &pb.UnreadyRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("unready response failed, err: %v", err)
		}
	}()

	player, err := l.playerMgr.GetPlayer(ctx.Request.UID)
	if err != nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	seat := player.Seat()
	if seat == nil {
		res.Code = common.Code_IllegalOperation
		return
	}

	if seat.IsUnready() {
		res.Code = common.Code_OK
		return
	}

	seat.Unready()

	l.syncSeatStateChange(seat, pb.SeatState_Unready)

	res.Code = common.Code_OK
}

// 同步游戏信息
func (l *Mahjong) syncGameInfo(player *entity.Player) {
	err := l.proxy.Push(l.ctx, &cluster.PushArgs{
		Kind:   session.User,
		Target: player.UID(),
		Message: &cluster.Message{
			Route: route.GameInfoNotify,
			Data: &pb.GameInfoNotify{
				GameInfo: l.makeGameInfo(player),
			},
		},
	})
	if err != nil {
		log.Errorf("game info notify push failed: %v", err)
	}
}

// 同步玩家座位状态
func (l *Mahjong) syncSeatStateChange(seat *entity.Seat, state pb.SeatState) {
	seats := seat.Table().Seats()
	targets := make([]int64, 0, len(seats))

	for _, s := range seats {
		if s.ID() == seat.ID() {
			continue
		}

		if s.IsOffline() {
			continue
		}

		p := s.Player()
		if p == nil {
			continue
		}

		targets = append(targets, p.UID())
	}

	if len(targets) == 0 {
		return
	}

	_, err := l.proxy.Multicast(l.ctx, &cluster.MulticastArgs{
		Kind:    session.User,
		Targets: targets,
		Message: &cluster.Message{
			Route: route.SeatStateChange,
			Data: &pb.SeatStateChangeNotify{
				SeatID:    int32(seat.ID()),
				SeatState: state,
			},
		},
	})
	if err != nil {
		log.Errorf("seat state change notify multicast failed: %v", err)
	}
}

// 同步玩家座位信息
func (l *Mahjong) syncTakeSeatInfo(seat *entity.Seat) {
	seats := seat.Table().Seats()
	targets := make([]int64, 0, len(seats))

	for _, s := range seats {
		if s.ID() == seat.ID() {
			continue
		}

		if s.IsOffline() {
			continue
		}

		p := s.Player()
		if p == nil {
			continue
		}

		targets = append(targets, p.UID())
	}

	if len(targets) == 0 {
		return
	}

	_, _ = l.proxy.Multicast(l.ctx, &cluster.MulticastArgs{
		Kind:    session.User,
		Targets: targets,
		Message: &cluster.Message{
			Route: route.TakeSeat,
			Data: &pb.TakeSeatNotify{
				Seat: &pb.Seat{
					ID:       int32(seat.ID()),
					IsOnline: seat.IsOnline(),
					IsReady:  seat.IsReady(),
					Player:   &pb.Player{User: l.makeUserInfo(seat.Player())},
				},
			},
		},
	})
}

// 根据玩家生成用户信息
func (l *Mahjong) makeUserInfo(player *entity.Player) *common.User {
	if player == nil {
		return nil
	}

	u := player.User()

	return &common.User{
		UID:           u.UID,
		Account:       u.Account,
		Nickname:      u.Nickname,
		Avatar:        u.Avatar,
		Signature:     u.Signature,
		Gender:        common.Gender(u.Gender),
		Level:         int32(u.Level),
		Experience:    int32(u.Experience),
		Coin:          int32(u.Coin),
		LastLoginIP:   u.LastLoginIP,
		LastLoginTime: u.LastLoginTime.Time().Unix(),
	}
}

// 根据玩家生成游戏信息
func (l *Mahjong) makeGameInfo(player *entity.Player) *pb.GameInfo {
	var (
		opts  = player.Room().Options()
		table = player.Table()
		seats = table.Seats()
	)

	info := &pb.GameInfo{
		Room: &pb.Room{
			ID:            int32(opts.ID),
			Name:          opts.Name,
			MinEntryLimit: int32(opts.MinEntryLimit),
			MaxEntryLimit: int32(opts.MaxEntryLimit),
		},
		Table: &pb.Table{
			ID:    int32(table.ID()),
			Seats: make([]*pb.Seat, len(seats)),
		},
	}

	for i, seat := range seats {
		info.Table.Seats[i] = &pb.Seat{
			ID:       int32(seat.ID()),
			IsOnline: seat.IsOnline(),
			IsReady:  seat.IsReady(),
		}

		p := seat.Player()
		if p == nil {
			continue
		}

		info.Table.Seats[i].Player = &pb.Player{
			IsMyself: player.UID() == p.UID(),
			User:     l.makeUserInfo(p),
		}
	}

	return info
}
