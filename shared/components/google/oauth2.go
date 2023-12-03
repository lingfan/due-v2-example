package google

import (
	"context"
	"github.com/dobyte/due/v2/config"
	"github.com/dobyte/due/v2/log"
	"google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"
	"sync"
)

var (
	oauth2Once     sync.Once
	oauth2Instance *oauth2.Service
)

func OAuth2() *oauth2.Service {
	oauth2Once.Do(func() {
		credentials := config.Get("config.google.oauth2.credentials").String()
		ins, err := oauth2.NewService(context.Background(), option.WithCredentialsFile(credentials))
		if err != nil {
			log.Fatalf("google oauth2 init failed: %v", err)
		}

		oauth2Instance = ins
	})

	return oauth2Instance
}
