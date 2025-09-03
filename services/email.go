package services

import (
	"backend/config"
	"fmt"
	"net/smtp"
)

func SendEmail(cfg config.Config, to, subject, body string) error {
	from := cfg.SMTPEmail
	pass := cfg.SMTPPassword

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}
