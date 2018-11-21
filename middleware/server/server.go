package main

import (
	"github.com/jhgv/gocodes/middleware/common/naming"
	"github.com/jhgv/gocodes/middleware/core/proxy"
	"github.com/jhgv/gocodes/middleware/server/invoker"
)

const (
	host       = "localhost"
	port       = 7777
	namingServerPort       = 9995
	objectName = "FileConverter"
	objectID   = 9999
)

func main() {

	namingservice := naming.NewNamingProxy(host, namingServerPort)
	fileConverterInvoker := &invoker.FileConverterInvoker{}
	fileConverterProxy := proxy.FileConverterProxy{}
	fileConverterProxy.SetID(objectID)
	fileConverterProxy.SetHost(host)
	fileConverterProxy.SetPort(port)

	go namingservice.Bind(objectName, fileConverterProxy)
	fileConverterInvoker.Invoke(fileConverterProxy)
}
