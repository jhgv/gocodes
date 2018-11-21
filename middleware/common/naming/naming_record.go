package naming

import "github.com/jhgv/gocodes/middleware/core/proxy"

type NamingRecord struct {
	serviceName string
	clientProxy proxy.FileConverterProxy
}

func (nr *NamingRecord) GetServiceName() string{
	return nr.serviceName
}

func (nr *NamingRecord) SetServiceName(serviceName string) {
	nr.serviceName = serviceName
}

func (nr *NamingRecord) GetClientProxy() proxy.FileConverterProxy{
	return nr.clientProxy
}

func (nr *NamingRecord) SetClientProxy(clientProxy proxy.FileConverterProxy) {
	nr.clientProxy = clientProxy
}


