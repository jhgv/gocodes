package handler

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/jhgv/gocodes/middleware/patterns/utils/protocols"

	"github.com/jhgv/gocodes/middleware/patterns/upperfy"
)

type RPCServerRequestHanlder struct {
	conn     net.Conn
	listener net.Listener
	textfy   *upperfy.Textfy
}

func (ch *RPCServerRequestHanlder) SetupSocket(host string, port int) error {
	address := fmt.Sprintf(":%d", port)
	ch.textfy = new(upperfy.Textfy)
	rpc.Register(ch.textfy)
	rpc.HandleHTTP()
	listener, err := net.Listen(protocols.TCP, address)
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	log.Printf("Server listenning on %s ...\n", address)
	// No handler
	http.Serve(listener, nil)
	return nil
}

func (ch *RPCServerRequestHanlder) Send(message []byte) error {
	return nil
}

func (ch *RPCServerRequestHanlder) Recieve() ([]byte, error) {

	return nil, nil
}
