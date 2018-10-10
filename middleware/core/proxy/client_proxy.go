package proxy

type ClientProxy struct {
	host string
	port int
	id   int
}

func (cp *ClientProxy) GetHost() string {
	return cp.host
}
func (cp *ClientProxy) SetHost(host string) {
	cp.host = host
}
func (cp *ClientProxy) GetPort() int {
	return cp.port
}
func (cp *ClientProxy) SetPort(port int) {
	cp.port = port
}

func (cp *ClientProxy) SetID(id int) {
	cp.id = id
}

func (cp *ClientProxy) GetID() int {
	return cp.id
}
