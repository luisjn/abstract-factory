package main

import "fmt"

type iSender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type iNotification interface {
	SendNotification()
	GetSender() iSender
}

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

func getNotificationFactory(notificationType string) (iNotification, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("no iNotification type")
}

func sendNotification(n iNotification) {
	n.SendNotification()
}

func getMethod(n iNotification) {
	fmt.Println(n.GetSender().GetSenderMethod())
}
