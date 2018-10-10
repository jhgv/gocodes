package main

import (
	"github.com/jhgv/gocodes/middleware/core/proxy"
	"github.com/jhgv/gocodes/middleware/server/invoker"
)

const (
	host       = "localhost"
	port       = 8081
	objectName = "TextHelper"
	objectID   = 9999
)

func main() {
	invoker := &invoker.TextHelperInvoker{}
	textHelperProxy := &proxy.TextHelperProxy{}
	textHelperProxy.SetID(objectID)
	textHelperProxy.SetHost(host)
	textHelperProxy.SetPort(port)

	invoker.Invoke(textHelperProxy)
}
