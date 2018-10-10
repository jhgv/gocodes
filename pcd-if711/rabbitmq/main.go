package main

import (
	"github.com/jhgv/gocodes/middleware/rabbitmq/client"
	"github.com/jhgv/gocodes/middleware/rabbitmq/server"
)

func main() {
	forever := make(chan bool)
	go client.StartClient()
	go server.StartServer()
	<- forever
}