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
	fromEmail := helpers.GetEnv("HOST_EMAIL", "")
	host := helpers.GetEnv("HOST", "")
	port := helpers.GetEnv("PORT", "")
	user := helpers.GetEnv("USER_KEY", "")
	pass := helpers.GetEnv("PASS_KEY", "")

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
