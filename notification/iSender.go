package notification

type iSender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}
