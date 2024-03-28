package utils

import (
	"net/smtp"
)

type mail struct {
	SMTP struct {
		Host string
		PORT string
	}
	Password string
	From     string
	To       []string
	Message  []byte
}

func NewMail(host, port, username, password string) *mail {
	newMail := new(mail)
	newMail.SMTP.Host = host
	newMail.SMTP.PORT = port
	newMail.From = username
	newMail.Password = password
	return newMail
}

func (m *mail) address() string {
	return m.SMTP.Host + ":" + m.SMTP.PORT
}

func (m *mail) authentication() smtp.Auth {
	return smtp.PlainAuth("", m.From, m.Password, m.SMTP.Host)
}

func (m *mail) SendMail(receivers []string, subject, message string) error {
	m.To = receivers

	to := ""
	for _, address := range receivers {
		to += address + ", "
	}

	m.Message = []byte("To: " + to + "\r\n" + "Subject: " + subject + "\r\n\r\n" + message)
	return smtp.SendMail(m.address(), m.authentication(), m.From, m.To, m.Message)
}
