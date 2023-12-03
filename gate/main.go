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
	"github.com/dobyte/due/transport/grpc/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/gate"
	"github.com/dobyte/due/v2/mode"
	"net/http"
)

func main() {
	mode.SetMode(mode.DebugMode)

	// 创建容器
	container := due.NewContainer()
	// 创建服务器
	server := ws.NewServer()
	// 监听HTTP连接升级WS协议
	server.OnUpgrade(func(w http.ResponseWriter, r *http.Request) bool {
		return true
	})
	// 创建网关组件
	var component = gate.NewGate(
		gate.WithServer(server),
		gate.WithLocator(redis.NewLocator()),
		gate.WithRegistry(etcd.NewRegistry()),
		gate.WithTransporter(grpc.NewTransporter()),
	) // 添加网关组件
	container.Add(component)
	// 启动容器
	container.Serve()
}
