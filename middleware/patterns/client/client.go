package main

import (
	"log"
	"time"

	"github.com/jhgv/gocodes/middleware/patterns/client/handler"
	"github.com/jhgv/gocodes/middleware/patterns/utils/constants"
	"github.com/jhgv/gocodes/middleware/patterns/utils/protocols"
	"github.com/jhgv/gocodes/middleware/utils"
)

const (
	protocol = "rpc"
	host     = "localhost"
	port     = 8081
)

func StartClient(client handler.ClientRequestHandler) {
	totalTime := 0.0
	// xlsBuilder := utils.XlsxBuilder{}
	// xlsBuilder.SetBasicMetricsFile(protocol, constants.NumRepetitions)
	for i := 0; i < constants.NumRepetitions; i++ {
		client.SetupSocket(host, port)
		// if err != nil {
		// 	log.Fatal("Error starting connection: ", err)
		// }
		message := utils.GenerateRandomText(200)

		startReq := time.Now()
		client.Send([]byte(message))
		// if err != nil {
		// 	log.Fatal("Error starting connection: ", err)
		// }
		client.Recieve()
		elapsedReq := time.Since(startReq)
		totalTime = totalTime + (elapsedReq.Seconds() * 1000.0)
		// log.Printf("%f", elapsedReq.Seconds()*1000.0)
		//xlsBuilder.AddRowData(elapsedReq.Seconds() * 1000.0)
		time.Sleep(10 * time.Millisecond)
		// if err != nil {
		// 	log.Fatal("Error recieveing message from server: ", err)
		// }
		// log.Printf("Message from server: { %s }\n", string(messageFromServer))
	}
	log.Printf("Average request time: %f", totalTime/float64(constants.NumRepetitions))
	//xlsBuilder.GenerateFile()
}

func main() {
	switch protocol {
	case protocols.TCP:
		client := new(handler.TCPClientHandler)
		StartClient(client)
	case protocols.UDP:
		client := new(handler.UDPClientHandler)
		StartClient(client)
	case protocols.RPC:
		client := new(handler.RPCClientHandler)
		StartClient(client)
	default:
		log.Println("Not available yet")
	}
}
