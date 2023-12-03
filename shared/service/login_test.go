package service_test

import (
	"due-v2-example/shared/service"
	"testing"
)

func TestLogin_GuestLogin(t *testing.T) {
	svc := service.NewLogin(nil)

	uid, err := svc.GuestLogin("1", "127.0.0.1")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(uid)
}
