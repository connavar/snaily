package email

import (
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/model"
	"gopkg.in/gomail.v2"
)

import (
	"crypto/tls"
)

var GlobalMailConfig model.OutgoingMailConfig

func GetDialer(config model.OutgoingMailConfig) *gomail.Dialer {
	d := gomail.NewDialer(config.SMTPHost, config.Port, config.Username, config.Password)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         "smtp.migadu.com",
	}
	return d

}

func send(dialer *gomail.Dialer, m ...*gomail.Message) (err error) {
	if err = dialer.DialAndSend(m...); err != nil {
		logrus.Errorln(err)
	}
	return
}

func SendMail(mail *model.Mail) (err error) {
	return send(GetDialer(GlobalMailConfig), mailToMessage(mail))
}

func SendMailWithConfig(mail *model.Mail, config model.OutgoingMailConfig) (err error) {
	return send(GetDialer(config), mailToMessage(mail))
}

func mailToMessage(mail *model.Mail) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", mail.FromEmail)
	m.SetHeader("To", mail.ToEmail)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.BodyHTML)
	return m

}
