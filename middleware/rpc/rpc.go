package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/jhgv/gocodes/middleware/patterns/upperfy"
	"github.com/jhgv/gocodes/middleware/patterns/utils/protocols"
	"github.com/jhgv/gocodes/middleware/utils"
	"github.com/jhgv/gocodes/middleware/utils/constants"
)

const protocol = "rpc"

func startClient(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	xlsBuilder := utils.XlsxBuilder{}
	xlsBuilder.SetBasicMetricsFile(protocol, constants.NumRepetitions)
	for i := 0; i < constants.NumRepetitions; i++ {
		client, _ := rpc.DialHTTP("tcp", address)
		// if err != nil {
		// 	log.Fatal("Error starting connection: ", err)
		// }
		message := utils.GenerateRandomText(200)
		startReq := time.Now()
		args := &upperfy.Args{Text: message}
		var reply string
		client.Call("Textfy.UpperText", args, &reply)
		// if err != nil {
		// 	log.Fatal("Error starting connection: ", err)
		// }
		elapsedReq := time.Since(startReq)
		// log.Printf("%f", elapsedReq.Seconds()*1000.0)
		xlsBuilder.AddRowData(elapsedReq.Seconds() * 1000.0)
		time.Sleep(10 * time.Millisecond)
		// if err != nil {
		// 	log.Fatal("Error recieveing message from server: ", err)
		// }
		// log.Printf("Message from server: { %s }\n", string(messageFromServer))
	}
}

func startServer(port int) {
	address := fmt.Sprintf(":%d", port)
	textfy := new(upperfy.Textfy)
	rpc.Register(textfy)
	rpc.HandleHTTP()
	listener, err := net.Listen(protocols.TCP, address)
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	log.Printf("Server listenning on %s ...\n", address)
	// No handler
	http.Serve(listener, nil)
}

func main() {

}
