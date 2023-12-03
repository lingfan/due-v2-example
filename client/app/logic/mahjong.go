package logic

import (
	"due-v2-example/shared/pb/common"
	mahjongpb "due-v2-example/shared/pb/mahjong"
	"due-v2-example/shared/route"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/log"
)

type Mahjong struct {
	proxy *client.Proxy
}

func NewMahjong(proxy *client.Proxy) *Mahjong {
	return &Mahjong{proxy: proxy}
}

func (l *Mahjong) Init() {
	// 快速开始回执
	l.proxy.AddRouteHandler(route.QuickStart, l.quickStartAck)
	// 游戏信息通知
	l.proxy.AddRouteHandler(route.GameInfoNotify, l.gameInfoNotify)
}

func (l *Mahjong) quickStartAck(ctx *client.Context) {
	res := &mahjongpb.QuickStartRes{}

	err := ctx.Parse(res)
	if err != nil {
		log.Errorf("invalid quick start ack message, err: %v", err)
		return
	}

	if res.Code != common.Code_OK {
		log.Warnf("%v", res.Code)
		return
	}

	log.Info("quick start success")
	log.Infof("%+v", res)
}

func (l *Mahjong) gameInfoNotify(ctx *client.Context) {
	ntf := &mahjongpb.GameInfoNotify{}

	err := ctx.Parse(ntf)
	if err != nil {
		log.Errorf("invalid game info notify message, err: %v", err)
		return
	}

	log.Info("receive game info")
	log.Infof("%+v", ntf)
}
