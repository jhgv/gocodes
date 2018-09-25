package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"

	"github.com/jhgv/gocodes/middleware/rpc/upperfy"
	"github.com/jhgv/gocodes/middleware/utils"
)

const numRepetitions int = 1000
const serverPort string = "1234"
const serverHost string = "localhost"

func main() {
	serverAdress := fmt.Sprintf("%s:%s", serverHost, serverPort)
	client, err := rpc.DialHTTP("tcp", serverAdress)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	xlsBuilder := utils.XlsxBuilder{}
	fileName := fmt.Sprintf("rpc-%d.xlsx", numRepetitions)
	xlsBuilder.SetBasicMetricsFile(fileName, numRepetitions)

	for i := 0; i < numRepetitions; i++ {
		text := utils.GenerateRandomText(30)
		args := &upperfy.Args{Text: text}
		var reply string
		start := time.Now()
		err = client.Call("Textfy.UpperText", args, &reply)
		elapsed := time.Since(start)
		xlsBuilder.AddRowData(elapsed.Seconds() * 1000)
		if err != nil {
			log.Fatal("textfy error: ", err)
		}
		// log.Printf("Changed from { %s } to { %s }", text, reply)

		time.Sleep(10 * time.Millisecond)
	}
	xlsBuilder.GenerateFile()

}
