package service

import (
	"context"
	"due-v2-example/shared/code"
	"due-v2-example/shared/components/mongo"
	userdao "due-v2-example/shared/dao/user"
	usermodel "due-v2-example/shared/model/user"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/errors"
)

type User struct {
	ctx     context.Context
	proxy   *node.Proxy
	userDao *userdao.User
}

func NewUser(proxy *node.Proxy) *User {
	return &User{
		ctx:     context.Background(),
		proxy:   proxy,
		userDao: userdao.NewUser(mongo.DB()),
	}
}

// GetUser 获取用户
// code.NotFoundUser
// code.InternalServerError
func (s *User) GetUser(uid int64) (*usermodel.User, error) {
	user, err := s.userDao.FindOneByUID(s.ctx, uid)
	if err != nil {
		return nil, errors.NewError(code.InternalServerError, err)
	}

	if user == nil {
		return nil, errors.NewError(code.NotFoundUser)
	}

	return user, nil
}
