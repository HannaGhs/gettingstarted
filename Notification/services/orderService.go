package services

import (
		"notification/entities"
		"notification/externalServices"
)

type OrderService struct {
	Notifier externalservices.Notifier
}

func (orderService *OrderService) CreateOrder(order *entities.Order) *entities.Order {
	orderService.Notifier = externalservices.NewNotifier(order.NotificationType)
	orderService.Notifier.SendNotify(order.UserID, "Order created")

	return order
}

func NewOrderService() *OrderService {
	return &OrderService{}
}