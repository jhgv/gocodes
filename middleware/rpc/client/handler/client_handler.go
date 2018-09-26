package handler

// ClientRequestHandler : Interface for client request handlers
type ClientRequestHandler interface {
	SetupSocket(host string, port int) error
	Send(message []byte) error
	Recieve() ([]byte, error)
}
