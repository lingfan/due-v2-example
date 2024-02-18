package logic

import (
	"due-v2-example/client/app/store"
	commonpb "due-v2-example/shared/pb/common"
	loginpb "due-v2-example/shared/pb/login"
	"due-v2-example/shared/route"

	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/log"
)

type Login struct {
	proxy *client.Proxy
}

func NewLogin(proxy *client.Proxy) *Login {
	return &Login{proxy: proxy}
}

func (l *Login) Init() {
	// 用户注册
	l.proxy.AddRouteHandler(route.Register, l.register)
	// 用户登录
	l.proxy.AddRouteHandler(route.Login, l.login)
}

// 用户注册
func (l *Login) register(ctx *client.Context) {

}

// 用户登录
func (l *Login) login(ctx *client.Context) {
	res := &loginpb.LoginRes{}

	err := ctx.Parse(res)
	if err != nil {
		log.Errorf("invalid login message, err: %v", err)
		return
	}

	if res.Code != commonpb.Code_OK {
		log.Warnf("logic.login %+v", res)
		return
	}

	store.Token = res.Token

	log.Info("login success")

	msg := &cluster.Message{
		Route: route.QuickStart,
		Data:  []byte{},
	}
	err = ctx.Conn().Push(msg)
	if err != nil {
		log.Errorf("push route.QuickStart message failed: %v", err)
	}
}
