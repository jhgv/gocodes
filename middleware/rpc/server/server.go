package main

import (
	"log"
	"strings"

	"github.com/jhgv/gocodes/middleware/rpc/server/handler"
	"github.com/jhgv/gocodes/middleware/rpc/utils/protocols"
)

const (
	protocol = "rpc"
	host     = "localhost"
	port     = 8081
)

func StartServer(server handler.ServerRequestHandler) {
	server.SetupSocket(host, port)
	// if err != nil {
	// 	log.Fatal("Error setting socket up", err)
	// }
	for {
		message, _ := server.Recieve()
		// if err != nil {
		// 	log.Fatal("Error recieving message from client", err)
		// }
		server.Send(message)
	}

}

func main() {
	forever := make(chan bool)
	log.Printf("Starting %s handler server\n", strings.ToUpper(protocol))
	switch protocol {
	case protocols.TCP:
		server := new(handler.TCPServerRequestHanlder)
		go StartServer(server)
	case protocols.UDP:
		server := new(handler.UDPServerRequestHanlder)
		go StartServer(server)
	case protocols.RPC:
		server := new(handler.RPCServerRequestHanlder)
		StartServer(server)
	default:
		log.Println("Not available yet")
	}
	<-forever
}
