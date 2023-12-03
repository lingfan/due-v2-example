package service

import (
	"context"
	"due-v2-example/shared/code"
	"due-v2-example/shared/components/google"
	"due-v2-example/shared/components/jwt"
	"due-v2-example/shared/components/mongo"
	userdao "due-v2-example/shared/dao/user"
	usermodel "due-v2-example/shared/model/user"
	"due-v2-example/shared/utils/xcrypt"
	"due-v2-example/shared/utils/xvalidate"
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xconv"
	jwtpkg "github.com/dobyte/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Login struct {
	ctx     context.Context
	proxy   *node.Proxy
	jwt     *jwt.JWT
	userDao *userdao.User
}

func NewLogin(proxy *node.Proxy) *Login {
	return &Login{
		ctx:     context.Background(),
		proxy:   proxy,
		jwt:     jwt.Instance(),
		userDao: userdao.NewUser(mongo.DB()),
	}
}

// GuestLogin 来宾登录
// code.InternalServerError
func (s *Login) GuestLogin(deviceID string, clientIP string) (int64, error) {
	user, err := s.userDao.FindOneByDeviceID(s.ctx, usermodel.TypeGuest, deviceID)
	if err != nil {
		return 0, errors.NewError(code.InternalServerError, err)
	}

	if user == nil {
		user = &usermodel.User{
			Type:        usermodel.TypeGuest,
			DeviceID:    deviceID,
			RegisterIP:  clientIP,
			LastLoginIP: clientIP,
		}
		if err = s.doCreateUser(user); err != nil {
			return 0, err
		}
	} else {
		s.doUpdateLoginRecord(user.ID, clientIP)
	}

	return user.UID, nil
}

// MobileLogin 手机号登录
// code.InternalServerError
func (s *Login) MobileLogin(mobile string, captcha string, clientIP string) (int64, error) {
	user, err := s.userDao.FindOneByMobile(s.ctx, mobile)
	if err != nil {
		return 0, errors.NewError(code.InternalServerError, err)
	}

	if user == nil {
		user = &usermodel.User{
			Type:        usermodel.TypeGeneral,
			Mobile:      mobile,
			RegisterIP:  clientIP,
			LastLoginIP: clientIP,
		}
		if err = s.doCreateUser(user); err != nil {
			return 0, err
		}
	} else {
		s.doUpdateLoginRecord(user.ID, clientIP)
	}

	return user.UID, nil
}

// AccountLogin 账号登录
// code.NotFoundUser
// code.WrongPassword
// code.InternalServerError
func (s *Login) AccountLogin(account string, password string, clientIP string) (int64, error) {
	var (
		err  error
		user *usermodel.User
	)

	switch {
	case xvalidate.IsEmail(account):
		user, err = s.userDao.FindOneByEmail(s.ctx, account)
	case xvalidate.IsMobile(account):
		user, err = s.userDao.FindOneByMobile(s.ctx, account)
	default:
		user, err = s.userDao.FindOneByAccount(s.ctx, account)
	}
	if err != nil {
		return 0, errors.NewError(code.InternalServerError, err)
	}

	if user == nil {
		return 0, errors.NewError(code.NotFoundUser)
	}

	ok, err := xcrypt.Compare(user.Password, password, user.Salt)
	if err != nil {
		return 0, errors.NewError(code.InternalServerError, err)
	}

	if !ok {
		return 0, errors.NewError(code.WrongPassword)
	}

	s.doUpdateLoginRecord(user.ID, clientIP)

	return user.UID, nil
}

// TokenLogin TOKEN登录
// code.TokenInvalid
// code.TokenExpired
func (s *Login) TokenLogin(token string, clientIP string) (int64, error) {
	identity, err := s.jwt.ExtractIdentity(token)
	if err != nil {
		switch {
		case jwtpkg.IsInvalidToken(err):
			return 0, errors.NewError(code.TokenInvalid, err)
		case jwtpkg.IsMissingToken(err):
			return 0, errors.NewError(code.TokenInvalid, err)
		case jwtpkg.IsExpiredToken(err):
			return 0, errors.NewError(code.TokenExpired, err)
		}
	}

	return xconv.Int64(identity), nil
}

// WechatLogin 微信登录
func (s *Login) WechatLogin(openid, deviceID, clientIP string) (int64, error) {
	return 0, nil
}

// GoogleLogin 谷歌登录
// code.InternalServerError
func (s *Login) GoogleLogin(idToken, deviceID, clientIP string) (int64, error) {
	tokenInfo, err := google.OAuth2().Tokeninfo().IdToken(idToken).Do()
	if err != nil {
		return 0, errors.NewError(code.InternalServerError, err)
	}

	user, err := s.userDao.FindOneByGoogleUserId(s.ctx, tokenInfo.UserId)
	if err != nil {
		return 0, errors.NewError(code.InternalServerError, err)
	}

	if user == nil {
		user = &usermodel.User{
			Type:           usermodel.TypeGeneral,
			ThirdPlatforms: usermodel.ThirdPlatforms{Google: tokenInfo.UserId},
			DeviceID:       deviceID,
			RegisterIP:     clientIP,
			LastLoginIP:    clientIP,
		}
		if err = s.doCreateUser(user); err != nil {
			return 0, err
		}
	} else {
		s.doUpdateLoginRecord(user.ID, clientIP)
	}

	return user.UID, nil
}

// 创建用户
// code.InternalServerError
func (s *Login) doCreateUser(user *usermodel.User) error {
	user.Coin = 100

	if _, err := s.userDao.InsertOne(s.ctx, user); err != nil {
		return errors.NewError(code.InternalServerError, err)
	}

	return nil
}

// 更新用户登录记录
func (s *Login) doUpdateLoginRecord(id primitive.ObjectID, clientIP string) {
	_, err := s.userDao.UpdateOne(s.ctx, func(cols *userdao.Columns) interface{} {
		return bson.M{cols.ID: id}
	}, func(cols *userdao.Columns) interface{} {
		return bson.M{"$set": bson.M{
			cols.LastLoginIP:   clientIP,
			cols.LastLoginTime: primitive.NewDateTimeFromTime(time.Now()),
		}}
	})
	if err != nil {
		log.Errorf("update user's login record failed, err: %v", err)
	}
}

// GenerateToken 生成Token
// code.InternalServerError
func (s *Login) GenerateToken(uid int64) (string, error) {
	token, err := s.jwt.GenerateToken(jwt.Payload{
		s.jwt.IdentityKey(): uid,
	})
	if err != nil {
		return "", errors.NewError(code.InternalServerError, err)
	}

	return token.Token, nil
}
