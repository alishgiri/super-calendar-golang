package services

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, body string) error {
	port := 587
	host := "smtp.gmail.com"

	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	fmt.Println(from)
	fmt.Println(password)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	message := fmt.Sprintf("From: Super Calendar (%s)\r\n", from)
	message += fmt.Sprintf("Subject: %s\n%s\r\n", subject, mime)
	message += fmt.Sprintf("\r\n%s\r\n", body)

	// Connect to the SMTP server.
	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", host, port), auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
