package event

import (
	"due-v2-example/client/app/store"
	pb "due-v2-example/shared/pb/login"
	"due-v2-example/shared/route"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/log"
	"math/rand"
)

func Init(proxy *client.Proxy) {
	proxy.AddEventListener(cluster.Connect, func(proxy *client.Proxy) {
		msg := &cluster.Message{
			Route: route.GameInfoNotify,
			Data: &pb.LoginReq{
				Mode:     pb.LoginMode_Guest,
				DeviceID: store.DeviceID + string(rune(rand.Intn(9))),
			},
		}

		err := proxy.Push(msg)
		if err != nil {
			log.Errorf("push message failed: %v", err)
		}

		for i := 0; i < 10; i++ {

			msg := &cluster.Message{
				Route: route.Login,
				Data: &pb.LoginReq{
					Mode:     pb.LoginMode_Guest,
					DeviceID: store.DeviceID + string(rune(i)),
				},
			}

			err := proxy.Push(msg)
			if err != nil {
				log.Errorf("push message failed: %v", err)
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
