package logic

import (
	"context"
	"due-v2-example/shared/code"
	"due-v2-example/shared/pb/common"
	pb "due-v2-example/shared/pb/login"
	"due-v2-example/shared/route"
	"due-v2-example/shared/service"
	"github.com/dobyte/due/v2/cluster"

	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/log"
)

type Login struct {
	proxy    *node.Proxy
	ctx      context.Context
	loginSvc *service.Login
}

func NewLogin(proxy *node.Proxy) *Login {
	return &Login{
		proxy:    proxy,
		ctx:      context.Background(),
		loginSvc: service.NewLogin(proxy),
	}
}

func (l *Login) Init() {
	l.proxy.Router().Group(func(group *node.RouterGroup) {
		// 用户注册
		group.AddRouteHandler(route.Register, false, l.register)
		// 用户登录
		group.AddRouteHandler(route.Login, false, l.login)
	})
}

// 用户注册
func (l *Login) register(ctx node.Context) {
	log.Infof("GetIP %v", ctx.CID())
}

// 用户登录
func (l *Login) login(ctx node.Context) {
	req := &pb.LoginReq{}
	res := &pb.LoginRes{}
	defer func() {
		if err := ctx.Response(res); err != nil {
			log.Errorf("login response failed, err: %v", err)
		}
	}()

	if err := ctx.Parse(req); err != nil {
		log.Errorf("invalid login message, err: %v", err)
		res.Code = common.Code_Abnormal
		return
	}

	clientIP, err := ctx.GetIP()
	if err != nil {
		log.Errorf("get client ip failed, err: %v", err)
		res.Code = common.Code_Abnormal
		return
	}
	log.Infof("GetIP %v", clientIP)

	if req.GetDeviceID() == "" {
		res.Code = common.Code_IllegalParams
		return
	}

	var uid int64
	switch req.GetMode() {
	case pb.LoginMode_Guest:
		uid, err = l.loginSvc.GuestLogin(req.GetDeviceID(), clientIP)
	case pb.LoginMode_Mobile:
		uid, err = l.loginSvc.MobileLogin(req.GetMobile(), req.GetCode(), clientIP)
	case pb.LoginMode_Account:
		uid, err = l.loginSvc.AccountLogin(req.GetAccount(), req.GetPassword(), clientIP)
	case pb.LoginMode_Wechat:
		uid, err = l.loginSvc.WechatLogin(req.GetOpenid(), req.GetDeviceID(), clientIP)
	case pb.LoginMode_Google:
		uid, err = l.loginSvc.GoogleLogin(req.GetToken(), req.GetDeviceID(), clientIP)
	case pb.LoginMode_Token:
		uid, err = l.loginSvc.TokenLogin(req.GetToken(), clientIP)
	default:
		log.Errorf("login failed, err: %v", err)
		res.Code = common.Code_IllegalParams
		return
	}
	if err != nil {
		switch errors.Code(err) {
		case code.NotFoundUser, code.WrongPassword:
			res.Code = common.Code_IncorrectAccountOrPassword
		case code.TokenExpired:
			res.Code = common.Code_TokenExpired
		default:
			res.Code = common.Code_Failed
		}
		log.Errorf("login failed, err: %v", err)
		return
	}

	if err = ctx.BindGate(uid); err != nil {
		log.Errorf("bind gate failed, err: %v", err)
		res.Code = common.Code_Failed
		return
	}

	token, err := l.loginSvc.GenerateToken(uid)
	if err != nil {
		log.Errorf("token generate failed, err: %v", err)
		res.Code = common.Code_Failed
		return
	}

	log.Infof("uid %v\ntoken %v\nnid %v\n", uid, token, ctx.NID)

	if err = l.proxy.BindNode(ctx.Context(), uid, "hall", "hall01"); err != nil {
		log.Errorf("bind node failed: uid: %d err: %v", ctx.UID, err)
		res.Code = common.Code_Failed
		return
	}

	nodes, err := l.proxy.FetchNodeList(ctx.Context(), cluster.Work)
	for _, _node := range nodes {
		log.Infof("FetchNodeList node:%#v", _node)
	}

	log.Infof("uid %v node name:%s, node id:%s", uid, l.proxy.GetName(), l.proxy.GetID())

	res.Code = common.Code_OK
	res.Token = token
}
