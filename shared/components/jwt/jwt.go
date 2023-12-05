package jwt

import (
	"github.com/dobyte/due/v2/etc"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/jwt"
	"sync"
)

var (
	once     sync.Once
	instance *jwt.JWT
)

type JWT = jwt.JWT
type Payload = jwt.Payload

// Instance 获取JWT实例
func Instance() *JWT {
	once.Do(func() {
		conf := &struct {
			Issuer        string `json:"issuer"`
			ValidDuration int    `json:"validDuration"`
			SecretKey     string `json:"secretKey"`
			IdentityKey   string `json:"identityKey"`
		}{}

		err := etc.Match("etc.jwt").Scan(conf)
		if err != nil {
			log.Fatalf("load jwt config failed: %v", err)
		}

		instance = jwt.NewJWT(
			jwt.WithIssuer(conf.Issuer),
			jwt.WithIdentityKey(conf.IdentityKey),
			jwt.WithSecretKey(conf.SecretKey),
			jwt.WithValidDuration(conf.ValidDuration),
		)
	})

	return instance
}
