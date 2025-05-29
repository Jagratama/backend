package resend

import (
	"fmt"
	"jagratama-backend/internal/config"

	"github.com/resend/resend-go/v2"
)

func SendMail(to, subject, body string) error {
	apiKey := config.GetEnv("RESEND_API_KEY", "")
	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "Jagratama <admin@jagratama.com>",
		To:      []string{to},
		Subject: subject,
		Html:    body,
	}
	_, err := client.Emails.Send(params)

	if err != nil {
		fmt.Printf("Failed to send email: %v", err)
		return err
	}

	return nil
}
