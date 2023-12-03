package xcrypt

import (
	"due-v2-example/shared/utils/xmd5"
	"golang.org/x/crypto/bcrypt"
)

// Encrypt 加密
func Encrypt(password, salt string) (string, error) {
	password = xmd5.MD5(xmd5.MD5(password) + salt)
	buf, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// Compare 对比秘钥
func Compare(hashed, password, salt string) (bool, error) {
	password = xmd5.MD5(xmd5.MD5(password) + salt)
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}

	return err == nil, err
}
