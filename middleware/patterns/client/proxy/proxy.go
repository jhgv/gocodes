package proxy

type ClientProxy struct {
	host string
	port int
}

func (cp *ClientProxy) getHost() string {
	return cp.host
}
func (cp *ClientProxy) setHost(host string) {
	cp.host = host
}
func (cp *ClientProxy) getPort() int {
	return cp.port
}
func (cp *ClientProxy) setPort(port int) {
	cp.port = port
}
