package service_test

import (
	"due-v2-example/shared/consts"
	"due-v2-example/shared/service"
	mailargs "due-v2-example/shared/service/args/mail"
	"testing"
)

var svc = service.NewMail(nil)

func TestMail_SendMail(t *testing.T) {
	mailID, err := svc.SendMail(6, mailargs.Sender{
		ID: -1,
	}, mailargs.Mail{
		Title:   "A Test Mail",
		Content: "test mail content",
		Attachments: []mailargs.Attachment{{
			PropID:  1,
			PropNum: 1,
		}},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(mailID)
}

func TestMail_ReadMail(t *testing.T) {
	err := svc.ReadMail("6385ac192be5c1b0ca23a08c", 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMail_ReadAllMail(t *testing.T) {
	err := svc.ReadAllMail(consts.Administrator)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMail_DeleteMail(t *testing.T) {
	err := svc.DeleteMail("637e16d788457148afd53496", consts.Administrator, true)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMail_DeleteAllMail(t *testing.T) {
	err := svc.DeleteAllMail(consts.Administrator, false)
	if err != nil {
		t.Fatal(err)
	}
}
