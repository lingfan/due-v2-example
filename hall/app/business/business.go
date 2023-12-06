package business

import (
	"context"
	"due-v2-example/hall/app/logic"
	"due-v2-example/shared/event"

	"github.com/dobyte/due/eventbus/nats/v2"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/eventbus"
	"github.com/dobyte/due/v2/log"
)

func Init(proxy *node.Proxy) {
	// 初始化路由
	initRoute(proxy)
	// 初始化事件
	initEvent()
}

// 初始化路由
func initRoute(proxy *node.Proxy) {
	// 登录逻辑
	logic.NewLogin(proxy).Init()
	// 邮件逻辑
	logic.NewMail(proxy).Init()
}

// 初始化事件
func initEvent() {
	// 初始化事件总线
	eventbus.SetEventbus(nats.NewEventbus())
	// 订阅用户登录事件
	err := eventbus.Subscribe(context.Background(), event.Login, event.LoginHandler)
	if err != nil {
		log.Fatalf("%s event subscribe failed: %v", event.Login, err)
	}

	err = eventbus.Publish(context.Background(), event.Login, "login")
	if err != nil {
		log.Fatalf("%s event subscribe failed: %v", event.Login, err)
	}
}
