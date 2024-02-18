package main

import (
	"due-v2-example/game/app/business"
	"github.com/dobyte/due/transport/rpcx/v2"

	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/etcd/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/mode"
)

func main() {
	mode.SetMode(mode.DebugMode)

	// 创建容器
	container := due.NewContainer()
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := etcd.NewRegistry()
	// 创建RPC传输器
	transporter := rpcx.NewTransporter()
	// 创建网关组件user
	component := node.NewNode(
		node.WithLocator(locator),
		node.WithRegistry(registry),
		node.WithTransporter(transporter),
	)
	// 初始化业务
	business.Init(component.Proxy())
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
