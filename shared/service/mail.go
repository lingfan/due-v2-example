package service

import (
	"context"
	"due-v2-example/shared/code"
	"due-v2-example/shared/components/mongo"
	"due-v2-example/shared/consts"
	"due-v2-example/shared/route"
	"github.com/dobyte/due/v2/cluster"

	maildao "due-v2-example/shared/dao/mail"
	mailmodel "due-v2-example/shared/model/mail"
	mailpb "due-v2-example/shared/pb/mail"
	mailargs "due-v2-example/shared/service/args/mail"

	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/session"
	"go.mongodb.org/mongo-driver/bson"
)

type Mail struct {
	ctx     context.Context
	proxy   *node.Proxy
	mailDao *maildao.Mail
	userSvc *User
}

func NewMail(proxy *node.Proxy) *Mail {
	return &Mail{
		ctx:     context.Background(),
		proxy:   proxy,
		mailDao: maildao.NewMail(mongo.DB()),
		userSvc: NewUser(proxy),
	}
}

// SendMail 发送邮件
// code.NotFoundUser
// code.InternalServerError
func (s *Mail) SendMail(receiver int64, sender mailargs.Sender, mail mailargs.Mail) (string, error) {
	user, err := s.userSvc.GetUser(receiver)
	if err != nil {
		return "", errors.Replace(err, "not found receiver", *code.NotFoundUser)
	}

	if sender.ID > 0 {
		user, err = s.userSvc.GetUser(sender.ID)
		if err != nil {
			return "", errors.Replace(err, "not found sender", *code.NotFoundUser)
		}

		if sender.Name == "" {
			sender.Name = user.Nickname
		}

		if sender.Icon == "" {
			sender.Icon = user.Avatar
		}
	}

	model := &mailmodel.Mail{
		Title:       mail.Title,
		Content:     mail.Content,
		Receiver:    receiver,
		Status:      mailmodel.StatusUnread,
		Attachments: make([]mailmodel.Attachment, 0, len(mail.Attachments)),
		Sender: mailmodel.Sender{
			ID:   sender.ID,
			Name: sender.Name,
			Icon: sender.Icon,
		},
	}

	for _, attachment := range mail.Attachments {
		model.Attachments = append(model.Attachments, mailmodel.Attachment{
			PropID:  attachment.PropID,
			PropNum: attachment.PropNum,
		})
	}

	_, err = s.mailDao.InsertOne(s.ctx, model)
	if err != nil {
		return "", errors.NewError(code.InternalServerError, err)
	}

	if s.proxy != nil {
		data := &mailpb.Mail{
			ID:          model.ID.Hex(),
			Title:       model.Title,
			Content:     model.Content,
			Status:      mailpb.Status(model.Status),
			SendTime:    model.SendTime.Time().Unix(),
			Attachments: make([]*mailpb.Attachment, 0, len(mail.Attachments)),
			Sender: &mailpb.Sender{
				ID:   model.Sender.ID,
				Name: model.Sender.Name,
				Icon: model.Sender.Icon,
			},
		}

		for _, attachment := range mail.Attachments {
			data.Attachments = append(data.Attachments, &mailpb.Attachment{
				PropID:  int32(attachment.PropID),
				PropNum: int32(attachment.PropNum),
			})
		}

		_ = s.proxy.Push(s.ctx, &cluster.PushArgs{
			Kind:    session.User,
			Target:  receiver,
			Message: &cluster.Message{Route: route.NewMail, Data: data},
		})
	}

	return model.ID.String(), nil
}

// ReadMail 读取邮件
// code.NoPermission
// code.NotFoundMail
// code.InternalServerError
func (s *Mail) ReadMail(mailID string, owner int64) error {
	mail, err := s.GetMail(mailID)
	if err != nil {
		return err
	}

	if owner != consts.Administrator {
		if owner != mail.Receiver {
			return errors.NewError(code.NoPermission, "this mail does not belong to you")
		}

		if mail.Status != mailmodel.StatusUnread {
			return nil
		}
	}

	_, err = s.mailDao.UpdateOne(s.ctx, func(cols *maildao.Columns) interface{} {
		return bson.M{cols.ID: mail.ID, cols.Status: mailmodel.StatusUnread}
	}, func(cols *maildao.Columns) interface{} {
		return bson.M{"$set": bson.M{cols.Status: mailmodel.StatusRead}}
	})
	if err != nil {
		return errors.NewError(code.InternalServerError, err)
	}

	return nil
}

// ReadAllMail 读取所有邮件
func (s *Mail) ReadAllMail(owner int64) error {
	_, err := s.mailDao.UpdateMany(s.ctx, func(cols *maildao.Columns) interface{} {
		if owner != consts.Administrator {
			return bson.M{
				cols.Receiver: owner,
				cols.Status:   mailmodel.StatusUnread,
			}
		} else {
			return bson.M{cols.Status: mailmodel.StatusUnread}
		}
	}, func(cols *maildao.Columns) interface{} {
		return bson.M{"$set": bson.M{cols.Status: mailmodel.StatusRead}}
	})
	if err != nil {
		return errors.NewError(code.InternalServerError, err)
	}

	return nil
}

// DeleteMail 删除邮件
// code.NoPermission
// code.NotFoundMail
// code.InternalServerError
func (s *Mail) DeleteMail(mailID string, owner int64, isForce bool) error {
	mail, err := s.GetMail(mailID)
	if err != nil {
		return err
	}

	if owner != consts.Administrator {
		if owner != mail.Receiver {
			return errors.NewError(code.NoPermission, "this mail does not belong to you")
		}
	}

	if !isForce {
		if mail.Status == mailmodel.StatusUnread {
			return errors.NewError(code.NoPermission, "can't delete unread mail")
		}

		if len(mail.Attachments) > 0 && mail.Status != mailmodel.StatusReceived {
			return errors.NewError(code.NoPermission, "can't delete unreceived mail")
		}
	}

	_, err = s.mailDao.DeleteOneByID(s.ctx, mailID)
	if err != nil {
		return errors.NewError(code.InternalServerError, err)
	}

	return nil
}

// DeleteAllMail 删除所有邮件
// code.InternalServerError
func (s *Mail) DeleteAllMail(owner int64, isForce bool) error {
	_, err := s.mailDao.DeleteMany(s.ctx, func(cols *maildao.Columns) interface{} {
		filter := bson.M{}

		if owner != consts.Administrator {
			filter[cols.Receiver] = owner
		}

		if !isForce {
			filter["$or"] = bson.A{
				bson.M{cols.Status: mailmodel.StatusRead, cols.Attachments: bson.M{"$size": 0}},
				bson.M{cols.Status: mailmodel.StatusReceived},
			}
		}

		return filter
	})
	if err != nil {
		return errors.NewError(code.InternalServerError, err)
	}

	return nil
}

// GetMail 获取邮件
// code.NotFound
// code.InternalServerError
func (s *Mail) GetMail(mailID string) (*mailmodel.Mail, error) {
	mail, err := s.mailDao.FindOneByID(s.ctx, mailID)
	if err != nil {
		return nil, errors.NewError(code.InternalServerError, err)
	}

	if mail == nil {
		return nil, errors.NewError(code.NotFoundMail)
	}

	return mail, nil
}

func (s *Mail) List() {

}
