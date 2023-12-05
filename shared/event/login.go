package event

import (
	"github.com/dobyte/due/v2/eventbus"
	"github.com/dobyte/due/v2/log"
)

type LoginPayload struct {
	ID      int32
	Account string
}

// LoginHandler 登录事件处理器
func LoginHandler(event *eventbus.Event) {
	log.Infof("event LoginHandler %+v", event)
}
