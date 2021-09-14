package notification

import "fmt"

type INotificationFactory interface {
	CreateNotification() INotification
}

func GetNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("wrong Notification type passed")
}


