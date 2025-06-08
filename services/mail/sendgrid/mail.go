package sendgrid

import (
	"bytes"
	"fmt"
	"myturbogarage/helpers"
	"myturbogarage/services/mail"
	"net/smtp"
)

// SendEmail ...
func SendEmail(template, subject, to string, data map[string]interface{}) error {
	fromEmail := helpers.GetEnv("DEFAULT_FROM_EMAIL", "")
	host := helpers.GetEnv("EMAIL_HOST", "")
	port := helpers.GetEnvInt("EMAIL_PORT", 587)
	user := helpers.GetEnv("EMAIL_HOST_USER", "")
	pass := helpers.GetEnv("EMAIL_HOST_PASSWORD", "")

	data["title"] = subject
	htmlBody, err := mail.ParseTemplate(template, data)
	if err != nil {
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("From: AppInApp<%s>\n", fromEmail)))
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))
	body.Write([]byte(htmlBody))

	// Create an authentication mechanism for SMTP
	auth := smtp.PlainAuth("", user, pass, host)

	smtpHost := fmt.Sprintf("%s:%d", host, port)

	// Send the email using SMTP
	err = smtp.SendMail(smtpHost, auth, fromEmail, []string{to}, body.Bytes())
	if err != nil {
		return err
	}

	// Print a success message
	fmt.Println("Email sent successfully")

	return nil
}
