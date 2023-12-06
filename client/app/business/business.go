package business

import (
	"context"
	"due-v2-example/client/app/logic"
	"due-v2-example/client/app/store"
	"due-v2-example/shared/event"
	pb "due-v2-example/shared/pb/login"
	"due-v2-example/shared/route"
	"github.com/dobyte/due/eventbus/redis/v2"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/eventbus"
	"github.com/dobyte/due/v2/log"
)

func Init(proxy *client.Proxy) {

	// 初始化路由
	initRoute(proxy)
	// 初始化事件
	initEvent(proxy)
}

// 初始化路由
func initRoute(proxy *client.Proxy) {
	// 登录逻辑
	logic.NewLogin(proxy).Init()
	// 邮件逻辑
	logic.NewMail(proxy).Init()
	// 麻将逻辑
	logic.NewMahjong(proxy).Init()
}

// 初始化事件
func initEvent(proxy *client.Proxy) {

	// 初始化事件总线
	eventbus.SetEventbus(redis.NewEventbus())
	// 订阅用户登录事件
	err := eventbus.Subscribe(context.Background(), event.Login, event.LoginHandler)
	if err != nil {
		log.Fatalf("%s event subscribe failed: %v", event.Login, err)
	}

	err = eventbus.Publish(context.Background(), event.Login, "login test")
	if err != nil {
		log.Fatalf("%s event subscribe failed: %v", event.Login, err)
	}

	proxy.AddEventListener(cluster.Connect, func(proxy *client.Proxy) {

		for i := 0; i < 1; i++ {

			msg := &cluster.Message{
				Route: route.Login,
				Data: &pb.LoginReq{
					Mode:     pb.LoginMode_Guest,
					DeviceID: store.DeviceID + string(rune(i)),
				},
			}
			log.Infof("push route.Login message: %+v", msg)

			err := proxy.Push(msg)
			if err != nil {
				log.Errorf("push message route.Login failed: %v", err)
			}

		}

	})

	proxy.AddEventListener(cluster.Reconnect, func(proxy *client.Proxy) {
		if store.Token == "" {

			msg := &cluster.Message{
				Route: route.Login,
				Data: &pb.LoginReq{
					Mode:     pb.LoginMode_Guest,
					DeviceID: store.DeviceID,
				},
			}

			err := proxy.Push(msg)
			if err != nil {
				log.Errorf("push message failed: %v", err)
			}
		} else {
			msg := &cluster.Message{
				Route: route.Login,
				Data: &pb.LoginReq{
					Mode:     pb.LoginMode_Token,
					DeviceID: store.DeviceID,
					Token:    store.Token,
				},
			}

			err := proxy.Push(msg)
			if err != nil {
				log.Errorf("push message failed: %v", err)
			}
		}
	})

	//proxy.AddEventListener(cluster.Disconnect, func(proxy client.Proxy) {
	//	for {
	//		err := proxy.Reconnect()
	//		if err == nil {
	//			return
	//		}
	//
	//		time.Sleep(time.Second)
	//
	//		log.Errorf("reconnect failed: %v", err)
	//	}
	//})
}
