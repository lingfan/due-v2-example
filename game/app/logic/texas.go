package logic

import (
	"context"
	"due-v2-example/game/app/entity"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/config"
	"github.com/dobyte/due/v2/log"
)

type Texas struct {
	proxy     *node.Proxy
	ctx       context.Context
	roomMgr   *entity.RoomMgr
	playerMgr *entity.PlayerMgr
}

func NewTexas(proxy *node.Proxy) *Texas {
	opts := make([]*entity.Options, 0)
	err := config.Get("texas.rooms").Scan(&opts)
	if err != nil {
		log.Fatalf("load texas rooms config failed: %v", err)
	}

	return &Texas{
		proxy:     proxy,
		ctx:       context.Background(),
		roomMgr:   entity.NewRoomMgr(opts),
		playerMgr: entity.NewPlayerMgr(),
	}
}

func (l *Texas) Init() {

}
