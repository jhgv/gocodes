package naming

import (
	"github.com/jhgv/gocodes/middleware/core/proxy"
	"log"
)

type Naming struct{
	namingRepository *NamingRepository
}

func NewNaming() *Naming {
	n := new(Naming)
	n.namingRepository = &NamingRepository{}
	return n
}

func (n *Naming) Lookup(serviceName string) (proxy.FileConverterProxy, error) {
	//log.Printf("Looking for service %s\n", serviceName)
	var clientProxy proxy.FileConverterProxy
	clientProxy, err := n.namingRepository.getClientProxyByServiceName(serviceName)
	if err != nil {
		log.Print(err)
	}

	return clientProxy, nil
}

func (n *Naming) Bind(serviceName string, clientProxy proxy.FileConverterProxy) {
	log.Printf("Binding service %s\n", serviceName)
	n.namingRepository.addRecord(serviceName, clientProxy)
}
