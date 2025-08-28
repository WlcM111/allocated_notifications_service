package notifications_service

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailNotifier struct {
	APIKey string
	From   string
}

func CreateNewEmail(apiKey, from string) *EmailNotifier {
	return &EmailNotifier{APIKey: apiKey, From: from}
}

func (e *EmailNotifier) SendNotification(rp, item, message string) error {
	sn := mail.NewEmail("Notification Service", e.From)
	rcp := mail.NewEmail("User", rp)

	msg := mail.NewSingleEmail(sn, item, rcp, message, message)
	cl := sendgrid.NewSendClient(e.APIKey)

	rsp, err := cl.Send(msg)

	if err != nil {
		return err
	}

	if rsp.StatusCode >= 400 {
		log.Printf("sendgrid error: %v", rsp.Body)
		return fmt.Errorf("failed to send email")
	}

	log.Printf("Email sent to %s", rp)
	return nil
}
