package mail

import (
	"bytes"
	"html/template"
	"path/filepath"
	"app_food/pkg/config"
	"gopkg.in/mail.v2"
)


type Mailer struct {
	From     string
	Password string
	SMTPHost string
	SMTPPort int
}

var ML *Mailer 

func init() {
	from, _ := config.GetConfig("MAIL_FROM")
	password, _ := config.GetConfig("MAIL_PASSWORD")
	smtpHost, _ := config.GetConfig("MAIL_HOST")
	smtpPort := 587 

	ML = &Mailer{
		From:     from,
		Password: password,
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
	}
}

func (m *Mailer) SendEmail(to, subject, body string) error {
	message := mail.NewMessage()
	message.SetHeader("From", m.From)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	dialer := mail.NewDialer(m.SMTPHost, m.SMTPPort, m.From, m.Password)
	return dialer.DialAndSend(message)
}

func renderTemplate(templateName string, data interface{}) (string, error) {
	templatePath := filepath.Join("views", templateName)
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (m *Mailer) SendTemplateEmail(to, templateName string, subject string, data interface{}) error {
	body, err := renderTemplate(templateName, data)
	if err != nil {
		return err
	}

	return m.SendEmail(to, subject, body)
}
