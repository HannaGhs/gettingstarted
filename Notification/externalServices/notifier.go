package externalservices

import "notification/entities"

type Notifier interface {
	SendNotify(receiver string, message string)
}

func NewNotifier(notificationType entities.NotificationType) Notifier {
	switch notificationType {
	case entities.SMSNotification:
		return NewSmsService()
	case entities.EmailNotification:
		return NewEmailService()
	default:
		return NewNilNotifyService()
	}
}