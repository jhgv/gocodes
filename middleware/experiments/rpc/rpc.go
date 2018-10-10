package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
	"time"

	"github.com/jhgv/gocodes/middleware/experiments/rpc/application"
	"github.com/jhgv/gocodes/middleware/experiments/utils"
	"github.com/jhgv/gocodes/middleware/utils/constants"
	"github.com/jhgv/gocodes/middleware/utils/protocols"
)

const protocol = "rpc"
const host = "localhost"
const port int = 8081

func startClient(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	totalTime := 0.0

	for i := 0; i < constants.NumRepetitions; i++ {
		client, _ := rpc.DialHTTP("tcp", address)
		// if err != nil {
		// 	log.Fatal("Error starting connection: ", err)
		// }
		message := utils.GenerateRandomText(200)
		startReq := time.Now()
		args := &application.Args{Text: message}
		var reply string
		client.Call("Textfy.UpperText", args, &reply)
		// if err != nil {
		// 	log.Fatal("Error starting connection: ", err)
		// }
		elapsedReq := time.Since(startReq)
		totalTime = totalTime + (elapsedReq.Seconds() * 1000.0)
		// log.Printf("%f", elapsedReq.Seconds()*1000.0)
		// log.Printf("Message from server: { %s }\n", reply)
		time.Sleep(10 * time.Millisecond)

		// if err != nil {
		// 	log.Fatal("Error recieveing message from server: ", err)
		// }
	}
	log.Printf("Average request time: %f ms for %d requests", totalTime/float64(constants.NumRepetitions), constants.NumRepetitions)
}

func startServer(port int) {
	address := fmt.Sprintf(":%d", port)
	textfy := new(application.Textfy)
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
	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		startServer(port)
	} else {
		startClient(host, port)
	}
}
