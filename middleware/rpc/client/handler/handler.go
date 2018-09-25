package handler

// ClientRequestHandler : Interface for client request handlers
type ClientRequestHandler interface {
	SetupSockets() error
	Send() error
	Recieve() []byte
}
