package src

import (
	"fmt"

	"net/smtp"
	"os"
)

var emailAddress = os.Getenv("EMAIL_ADDRESS")
var emailPassword = os.Getenv("EMAIL_APP_PASSWORD")

// SendEmail sends an email using the app email address
func SendEmail(to string, subject string, message string) error {
	auth := smtp.PlainAuth("", emailAddress, emailPassword, "smtp.gmail.com")

	toSlice := []string{to}
	subject = fmt.Sprintf("Subject: %s\r\n", subject)
	fromLine := fmt.Sprintf("From: Notenheim <%s>\r\n", emailAddress)
	toLine := fmt.Sprintf("To: %s\r\n", to)
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	html := []byte(fromLine + toLine + subject + mime + message)

	err := smtp.SendMail("smtp.gmail.com:587", auth, emailAddress, toSlice, html)
	return err
}
