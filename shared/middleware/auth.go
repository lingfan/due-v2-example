package middleware

import (
	"due-v2-example/shared/route"
	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/log"
)

func Auth(middleware *node.Middleware, ctx node.Context) {
	log.Errorf("middleware Auth: %v", ctx.UID)

	if ctx.UID() == 0 {
		err := ctx.Reply(&cluster.Message{Route: route.Unauthorized})
		if err != nil {
			log.Errorf("response message failed, err: %v", err)
		}
	} else {
		middleware.Next(ctx)
	}
}
