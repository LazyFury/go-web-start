package util

import (
	"fmt"
	"net/smtp"
	"strings"
)

var (
	// Mail Mail
	Mail mail = mail{
		Host:     "smtp.qq.com",
		Nickname: "比鲁斯",
		Password: "dspbvexsvxmodiid",
		Port:     "587",
		User:     "2568597007@qq.com",
	}
)

type (
	mail struct {
		Nickname string
		User     string
		Password string
		Host     string
		Port     string
	}
)

func (m *mail) Auth() smtp.Auth {
	return smtp.PlainAuth("", m.User, m.Password, m.Host)
}

func (m *mail) SendMail(subject string, to []string, body string) (err error) {

	contentType := "Content-Type: text/html; charset=UTF-8"

	msg := []byte(fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s\r\n", strings.Join(to, ","), m.Nickname, m.User, subject, contentType, body))

	return smtp.SendMail(m.Host+":"+m.Port, m.Auth(), m.User, to, msg)
}
