package email

import (
	"github.com/team142/snaily/model"
	"os"
	"testing"
)

func TestSendMail(t *testing.T) {
	GlobalMailConfig = model.OutgoingMailConfig{
		SMTPHost: "smtp.migadu.com",
		Port:     587,
		Username: "new@dependmap.com",
		Password: os.Getenv("MAIL_PASSWORD"),
		UseTLS:   true,
	}

	mail := model.Mail{
		ToEmail:   "just1689@gmail.com",
		Subject:   "Testing migadu in Golang! Hope this works.",
		BodyHTML:  "<h1>Hi Justin</h1>Thats me!",
		FromEmail: "new@dependmap.com",
	}
	if err := SendMail(&mail); err != nil {
		t.Fail()
	}
}
