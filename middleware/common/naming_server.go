package main

import (
	"github.com/jhgv/gocodes/middleware/common/naming"
)

const (
	namingServerPort       = 9995
)

func main() {
	namingInvoker := naming.NewNamingInvoker()
	namingInvoker.Invoke(namingServerPort)
}
