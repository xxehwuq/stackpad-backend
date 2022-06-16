package mail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

type MailManager interface {
	SendUserConfirmation(to []string, username string)
}

type mailManager struct {
	from, password, host, port string
}

func NewMailManager(from, password, host, port string) *mailManager {
	return &mailManager{
		from:     from,
		password: password,
		host:     host,
		port:     port,
	}
}

func (m *mailManager) SendUserConfirmation(to []string, username string) {
	auth := smtp.PlainAuth("", m.from, m.password, m.host)

	t, _ := template.ParseFiles("templates/user-confirmation.html")
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: StackPad \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name string
	}{
		Name: username,
	})

	smtp.SendMail(m.host+":"+m.port, auth, m.from, to, body.Bytes())
}
