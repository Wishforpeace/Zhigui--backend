package mail

import (
	"gopkg.in/gomail.v2"
	_ "gopkg.in/gomail.v2"
)

var (
	mailer = gomail.NewDialer("smtp.qq.com", 465."1903180340@qq.com", "WYX&wyx20030310"
)

func Send(to, subject, body string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "your@qq.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if err = mailer.DialAndSend(m); err != nil {
		logger.LogError(err)
	}

	return err
}
