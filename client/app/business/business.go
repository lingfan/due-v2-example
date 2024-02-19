package business

import (
	"context"
	"due-v2-example/client/app/logic"
	"due-v2-example/client/app/store"
	rd "due-v2-example/shared/components/redis"
	"due-v2-example/shared/event"
	loginpb "due-v2-example/shared/pb/login"
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
	log.Infof("initRoute")

	// 登录逻辑
	logic.NewLogin(proxy).Init()
	// 邮件逻辑
	logic.NewMail(proxy).Init()
	// 麻将逻辑
	logic.NewMahjong(proxy).Init()
}

// 初始化事件
func initEvent(proxy *client.Proxy) {
	log.Infof("initEvent")

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

	// 监听组件启动
	proxy.AddHookListener(cluster.Start, startHandler)

	proxy.AddEventListener(cluster.Connect, connectHandler)

	proxy.AddEventListener(cluster.Reconnect, reconnectHandler)

	proxy.AddEventListener(cluster.Disconnect, disconnectHandler)
}

func disconnectHandler(conn *client.Conn) {
	log.Errorf("client disconnect: %v", conn.UID())

	//for {
	//reconnectHandler(conn)
	//time.Sleep(time.Second)
	//log.Errorf("reconnect failed: %v", err)
	//}
}

func reconnectHandler(conn *client.Conn) {
	if store.Token == "" {

		msg := &cluster.Message{
			Route: route.Login,
			Data: &loginpb.LoginReq{
				Mode:     loginpb.LoginMode_Guest,
				DeviceID: store.DeviceID,
			},
		}

		err := conn.Push(msg)
		if err != nil {
			log.Errorf("push message failed: %v", err)
		}
	} else {
		msg := &cluster.Message{
			Route: route.Login,
			Data: &loginpb.LoginReq{
				Mode:     loginpb.LoginMode_Token,
				DeviceID: store.DeviceID,
				Token:    store.Token,
			},
		}

		err := conn.Push(msg)
		if err != nil {
			log.Errorf("push message failed: %v", err)
		}
	}
}

// 组件启动处理器
func startHandler(proxy *client.Proxy) {
	if _, err := proxy.Dial(); err != nil {
		log.Errorf("client connect gate failed: %v", err)
		return
	}
}

// 连接建立处理器
func connectHandler(conn *client.Conn) {
	for i := 0; i < 1; i++ {

		msg := &cluster.Message{
			Route: route.Login,
			Data: &loginpb.LoginReq{
				Mode:     loginpb.LoginMode_Guest,
				DeviceID: store.DeviceID + string(rune(i)),
			},
		}

		//goredis 有bug ERR “Unknown subcommand or wrong number of arguments for 'setinfo'”
		//goredisRD := rd.GoRedisRD()
		//_key1 := "test1"
		//_ = goredisRD.Set(_key1, "test111")
		//_data := goredisRD.Get(_key1)
		//log.Infof("GoRedisRD.Get(%s): %+v", _key1, _data)

		RueidisRD := rd.RueidisRD()
		_key2 := "test2"
		_ = RueidisRD.Set(_key2, "test222")
		_data2 := RueidisRD.Get(_key2)
		log.Infof("RueidisRD.Get(%s): %+v", _key2, _data2)

		log.Infof("push route.Login message: %+v", msg)

		err := conn.Push(msg)
		if err != nil {
			log.Errorf("push message route.Login failed: %v", err)
		}

	}
}
