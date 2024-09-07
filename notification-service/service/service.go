package service

import (
	"gopkg.in/gomail.v2"
)

func SendMessage(to string, subject string, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "diyorilhomov676@gmail.com")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("Body", body)

	return nil
}