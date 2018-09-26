package main

import (
	"log"
	"time"

	"github.com/jhgv/gocodes/middleware/rpc/client/handler"
	"github.com/jhgv/gocodes/middleware/rpc/utils/constants"
	"github.com/jhgv/gocodes/middleware/rpc/utils/protocols"
	"github.com/jhgv/gocodes/middleware/utils"
)

const (
	protocol = "rpc"
	host     = "localhost"
	port     = 1234
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

		time.Sleep(10 * time.Millisecond)
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
		client := new(handler.RPCClientHandler)
		SendMessage(client)
	default:
		log.Println("Not available yet")
	}
}
