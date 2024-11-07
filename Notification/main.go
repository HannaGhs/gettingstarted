package main

import (
	"fmt"
	"notification/entities"
	"notification/services"
)

func main() {

	order := entities.Order{
		ID:           1,
		UserFullName: "John Doe",
		UserID:       "232456789",
		Price:        100.0,
		Status:       true,
		NotificationType: entities.SMSNotification,
	}

	order2 := entities.Order{
		ID:           2,
		UserFullName: "Jane Doe",
		UserID:       "Jane@gmail.com",
		Price:        100.0,
		Status:       true,
		NotificationType: entities.EmailNotification,
	}

	order3 := entities.Order{
		ID:           3,
		UserFullName: "Jack Smith",
		UserID:       "Jack@gmail.com",
		Price:        200.0,
		Status:       true,
		NotificationType: "",
	}

	orderService := services.NewOrderService()

	orderService.CreateOrder(&order)
	fmt.Printf("userID: %v notiftype: %v\n", order.UserID, order.NotificationType)

	orderService.CreateOrder(&order2)
	fmt.Printf("userID: %v notiftype: %v\n", order2.UserID, order2.NotificationType)

	orderService.CreateOrder(&order3)
	fmt.Printf("userID: %v notiftype: %v\n", order3.UserID, order3.NotificationType)
}