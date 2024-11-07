package externalservices

import (
	"fmt"
)

type SmsService struct {

}

func (e *SmsService) SendNotify(receiver string, message string) {
	fmt.Printf("SMS sent!\nreceiver: %s, message: %s\n", receiver, message)
}

func  NewSmsService() *SmsService {
	return &SmsService{}
}