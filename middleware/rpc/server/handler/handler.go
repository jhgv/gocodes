package handler

// ServerRequestHandler : Interface for server request handlers
type ServerRequestHandler interface {
	SetupSockets() error
	Send() error
	Recieve() []byte
}
