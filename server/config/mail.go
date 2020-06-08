package config

import (
	"fmt"
	"net/smtp"
	"strings"
)

type (
	// Mail 邮件配置
	mail struct {
		Nickname string `json:"nickname"`
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
	}
)

// Auth 身份认证
func (m *mail) Auth() smtp.Auth {
	return smtp.PlainAuth("", m.User, m.Password, m.Host)
}

// SendMail 发送邮件
func (m *mail) SendMail(subject string, to []string, body string) (err error) {

	contentType := "Content-Type: text/html; charset=UTF-8"

	msg := []byte(fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s\r\n", strings.Join(to, ","), m.Nickname, m.User, subject, contentType, body))

	return smtp.SendMail(m.Host+":"+m.Port, m.Auth(), m.User, to, msg)
}
