package handler

// ServerRequestHandler : Interface for server request handlers
type ServerRequestHandler interface {
	SetupSocket(host string, port int) error
	Send(message []byte) error
	Recieve() ([]byte, error)
}
