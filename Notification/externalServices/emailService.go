package externalservices

import (
	"fmt"
)

type EmailService struct {

}

func (e *EmailService) SendNotify(receiver string, message string) {
	fmt.Printf("Email sent!\nreceiver: %s, message: %s\n", receiver, message)
}

func NewEmailService() *EmailService {
	return &EmailService{}
}