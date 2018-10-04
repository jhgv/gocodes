package proxy

type ClientProxy interface {
	GetHost(host string, port int) error
	SetHost(message []byte) error
	GetPort() ([]byte, error)
	SetPort() ([]byte, error)
}
