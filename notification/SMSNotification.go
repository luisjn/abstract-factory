package notification

import "fmt"

type SMSNotificationSender struct{}
type SMSNotification struct{}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending iNotification via SMS...")
}

func (SMSNotification) GetSender() iSender {
	return SMSNotificationSender{}
}
