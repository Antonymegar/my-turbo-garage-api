package mail

import (
	"bytes"
	"fmt"
	"log"
	"myturbogarage/helpers"
	"net/smtp"
	tmp "text/template"
)

// SendEmail ...
func SendEmail(template, subject, to string, data map[string]interface{}) error {
	t, err := tmp.ParseFiles(template, "templates/base.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))

	data["title"] = subject
	if err := t.Execute(&body, data); err != nil {
		return err
	}

	from := helpers.GetEnv("DEFAULT_FROM_EMAIL", "")
	host := helpers.GetEnv("EMAIL_HOST", "")
	port := helpers.GetEnvInt("EMAIL_PORT", 587)
	user := helpers.GetEnv("EMAIL_HOST_USER", "")
	pass := helpers.GetEnv("EMAIL_HOST_PASSWORD", "")
	auth := smtp.PlainAuth("", user, pass, host)

	log.Printf("Sending email to %s, body: %s", to, body.String())

	smtpHost := fmt.Sprintf("%s:%d", host, port)
	err = smtp.SendMail(smtpHost, auth, from, []string{to}, body.Bytes())
	if err != nil {
		return err
	}

	return nil
}
