package main

import (
	"due-v2-example/client/app/business"
	"github.com/dobyte/due/network/ws/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/mode"
)

func main() {
	mode.SetMode(mode.DebugMode)

	// 创建容器
	container := due.NewContainer()
	// 创建网关组件
	component := client.NewClient(
		client.WithClient(ws.NewClient()),
	)
	// 初始化路由
	business.Init(component.Proxy())
	// 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
