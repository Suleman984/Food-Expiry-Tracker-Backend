package services

import (
	"backend/config"
	"fmt"
	"net/smtp"
)

func SendEmail(cfg config.Config, to, subject, body string) error {
	from := cfg.SMTPEmail
	pass := cfg.SMTPPassword

	msg := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
		"\r\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}
