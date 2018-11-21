package proxy

type ClientProxy struct {
	Host string
	Port int
	ID   int
}

func (cp *ClientProxy) GetHost() string {
	return cp.Host
}
func (cp *ClientProxy) SetHost(host string) {
	cp.Host = host
}
func (cp *ClientProxy) GetPort() int {
	return cp.Port
}
func (cp *ClientProxy) SetPort(port int) {
	cp.Port = port
}

func (cp *ClientProxy) SetID(id int) {
	cp.ID = id
}

func (cp *ClientProxy) GetID() int {
	return cp.ID
}
