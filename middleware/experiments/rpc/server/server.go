package server

import (
	"fmt"
	"github.com/jhgv/gocodes/middleware/experiments/rpc/application"
	"github.com/jhgv/gocodes/middleware/utils/protocols"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func StartServer(port int) {
	address := fmt.Sprintf(":%d", port)
	fileConverter := new(application.FileConverter)
	log.Print("Registering file converter object")
	rpc.Register(fileConverter)
	rpc.HandleHTTP()
	listener, err := net.Listen(protocols.TCP, address)
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	log.Printf("Server listenning on %s ...\n", address)
	// No handler
	http.Serve(listener, nil)
}
