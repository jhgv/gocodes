package main

import (
	"log"

	"github.com/jhgv/gocodes/middleware/rpc/client/handler"
	"github.com/jhgv/gocodes/middleware/rpc/utils/constants"
	"github.com/jhgv/gocodes/middleware/rpc/utils/protocols"
	"github.com/jhgv/gocodes/middleware/utils"
)

const (
	protocol = "udp"
	host     = "localhost"
	port     = 8081
)

func SendMessage(client handler.ClientRequestHandler) {

	for i := 0; i < constants.NumRepetitions; i++ {
		err := client.SetupSocket(host, port)
		if err != nil {
			log.Fatal("Error starting connection: ", err)
		}
		message := utils.GenerateRandomText(30)
		err = client.Send([]byte(message))
		if err != nil {
			log.Fatal("Error starting connection: ", err)
		}
		messageFromServer, err := client.Recieve()
		if err != nil {
			log.Fatal("Error recieveing message from server: ", err)
		}
		log.Printf("Message from server: { %s }\n", string(messageFromServer))
	}
}

func main() {
	switch protocol {
	case protocols.TCP:
		client := new(handler.TCPClientHandler)
		SendMessage(client)
	case protocols.UDP:
		client := new(handler.UDPClientHandler)
		SendMessage(client)
	case protocols.RPC:
		log.Println("Not available yet")
	default:
		log.Println("Not available yet")
	}
}
