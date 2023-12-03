package logic

import (
	pb "due-v2-example/shared/pb/mail"
	"due-v2-example/shared/route"
	"github.com/dobyte/due/v2/cluster/client"
	"github.com/dobyte/due/v2/log"
)

type Mail struct {
	proxy *client.Proxy
}

func NewMail(proxy *client.Proxy) *Mail {
	return &Mail{proxy: proxy}
}

func (l *Mail) Init() {
	// 新邮件
	l.proxy.AddRouteHandler(route.NewMail, l.newMail)
}

// 新邮件
func (l *Mail) newMail(ctx *client.Context) {
	res := &pb.Mail{}

	err := ctx.Parse(res)
	if err != nil {
		log.Errorf("invalid new mail message, err: %v", err)
		return
	}

	log.Infof("receive a new mail, %+v", res)
}
