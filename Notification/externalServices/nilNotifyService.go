package externalservices

import (
	"fmt"
)

type NilNotifyService struct {

}

func (e *NilNotifyService) SendNotify(receiver string, message string) {
	fmt.Printf("Please try again.\n")
}

func  NewNilNotifyService() *NilNotifyService {
	return &NilNotifyService{}
}