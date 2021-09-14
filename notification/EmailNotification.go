package notification

import "fmt"

type EmailNotificationSender struct{}
type EmailNotification struct{}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending iNotification via Email...")
}

func (EmailNotification) GetSender() iSender {
	return EmailNotificationSender{}
}
