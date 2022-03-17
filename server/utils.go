package server

type notificationEvent uint16

const (
	CLIENT_CONNECTED notificationEvent = iota
	CLIENT_DISCONNECTED
)

type notification struct {
	eventType  notificationEvent
	clientName string
}
