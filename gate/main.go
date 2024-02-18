/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2022/9/25 1:59 下午
 * @Desc: 网关服务器
 */

package main

import (
	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/network/ws/v2"
	"github.com/dobyte/due/registry/etcd/v2"
	"github.com/dobyte/due/transport/rpcx/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/gate"
	"github.com/dobyte/due/v2/mode"
)

func main() {
	mode.SetMode(mode.DebugMode)

	//这个自定义数据包 为uxgame 项目的解析数据包
	//uxPacker := net.NewPacker()
	//packet.SetPacker(uxPacker)

	// 创建容器
	container := due.NewContainer()
	// 创建服务器
	server := ws.NewServer()
	// 创建用户定位器
	locator := redis.NewLocator()
	// 创建服务发现
	registry := etcd.NewRegistry()
	// 创建RPC传输器
	transporter := rpcx.NewTransporter()
	// 创建网关组件
	var component = gate.NewGate(
		gate.WithServer(server),
		gate.WithLocator(locator),
		gate.WithRegistry(registry),
		gate.WithTransporter(transporter),
	) // 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
