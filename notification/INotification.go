package notification

type INotification interface {
	SendNotification()
	GetSender() iSender
}
