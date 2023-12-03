package middleware

import (
	"due-v2-example/shared/route"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/session"
)

func Auth(ctx *node.Context) {
	if ctx.Request.UID == 0 {
		err := ctx.Proxy.Push(ctx.Context(), &cluster.PushArgs{
			GID:     ctx.Request.GID,
			Kind:    session.Conn,
			Target:  ctx.Request.CID,
			Message: &cluster.Message{Route: route.Unauthorized},
		})
		if err != nil {
			log.Errorf("response message failed, err: %v", err)
		}
	} else {
		ctx.Middleware.Next(ctx)
	}
}
