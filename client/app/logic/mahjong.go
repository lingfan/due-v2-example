package logic

import (
	"due-v2-example/shared/pb/common"
	mahjongpb "due-v2-example/shared/pb/mahjong"
	mailpb "due-v2-example/shared/pb/mail"
	"due-v2-example/shared/route"
	"github.com/dobyte/due/v2/cluster"
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
	l.proxy.AddRouteHandler(route.Ready, l.ready)
	l.proxy.AddRouteHandler(route.ReadAllMail, l.readAllMail)

}

func (l *Mahjong) quickStartAck(ctx *client.Context) {
	res := &mahjongpb.QuickStartRes{}

	err := ctx.Parse(res)
	if err != nil {
		log.Errorf("invalid quick start ack message, err: %v", err)
		return
	}

	if res.Code != common.Code_OK {
		log.Warnf("quickStartAck %v", res.Code)
		return
	}

	log.Infof("quick start success %+v", res)

	msg := &cluster.Message{
		Route: route.Ready,
		Data:  []byte{},
	}
	err = l.proxy.Push(msg)
	if err != nil {
		log.Errorf("push message route.Ready failed: %v", err)
	}
}

func (l *Mahjong) gameInfoNotify(ctx *client.Context) {
	ntf := &mahjongpb.GameInfoNotify{}

	err := ctx.Parse(ntf)
	if err != nil {
		log.Errorf("invalid game info notify message, err: %v", err)
		return
	}

	log.Infof("receive game info %+", ntf)

	msg := &cluster.Message{
		Route: route.ReadAllMail,
		Data:  &mailpb.ReadAllMailReq{},
	}
	err = l.proxy.Push(msg)
	if err != nil {
		log.Errorf("push message route.FetchMailList failed: %v", err)
	}
	log.Infof("send route.ReadAllMail succ")

}

func (l *Mahjong) ready(ctx *client.Context) {
	log.Info("receive ready info")
}

func (l *Mahjong) readAllMail(ctx *client.Context) {
	res := &mailpb.ReadMailRes{}

	err := ctx.Parse(res)
	if err != nil {
		log.Errorf("invalid mailpb.ReadMailRes message, err: %v", err)
		return
	}

	log.Infof("receive ReadMailRes info %+v", res)
}
