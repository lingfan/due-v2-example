package user

import (
	"due-v2-example/shared/dao/user/internal"
	"due-v2-example/shared/model/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type Columns = internal.Columns

type User struct {
	*internal.User
}

func NewUser(db *mongo.Database) *User {
	return &User{User: internal.NewUser(db)}
}

// FindOneByUID 根据用户ID查找
func (dao *User) FindOneByUID(ctx context.Context, uid int64) (*user.User, error) {
	return dao.FindOne(ctx, func(cols *internal.Columns) interface{} {
		return bson.M{cols.UID: uid}
	})
}

// FindOneByMobile 根据手机号查找
func (dao *User) FindOneByMobile(ctx context.Context, mobile string) (*user.User, error) {
	return dao.FindOne(ctx, func(cols *internal.Columns) interface{} {
		return bson.M{cols.Mobile: mobile}
	})
}

// FindOneByEmail 根据邮箱查找
func (dao *User) FindOneByEmail(ctx context.Context, email string) (*user.User, error) {
	return dao.FindOne(ctx, func(cols *internal.Columns) interface{} {
		return bson.M{cols.Email: email}
	})
}

// FindOneByAccount 根据账号查找
func (dao *User) FindOneByAccount(ctx context.Context, account string) (*user.User, error) {
	return dao.FindOne(ctx, func(cols *internal.Columns) interface{} {
		return bson.M{cols.Account: account}
	})
}

// FindOneByDeviceID 根据设备ID查找
func (dao *User) FindOneByDeviceID(ctx context.Context, typ user.Type, deviceID string) (*user.User, error) {
	return dao.FindOne(ctx, func(cols *internal.Columns) interface{} {
		return bson.M{cols.Type: typ, cols.DeviceID: deviceID}
	})
}

// FindOneByGoogleUserId 根据谷歌用户ID查找
func (dao *User) FindOneByGoogleUserId(ctx context.Context, userId string) (*user.User, error) {
	return dao.FindOne(ctx, func(cols *internal.Columns) interface{} {
		return bson.M{cols.ThirdPlatforms: bson.M{"google": userId}}
	})
}
