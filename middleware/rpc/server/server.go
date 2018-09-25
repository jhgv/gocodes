package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/jhgv/gocodes/middleware/rpc/upperfy"
)

func main() {
	forever := make(chan int)
	textfy := new(upperfy.Textfy)
	rpc.Register(textfy)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	log.Println("Server listenning on port :8081 ...")
	// No handler
	go http.Serve(l, nil)
	<-forever
}
