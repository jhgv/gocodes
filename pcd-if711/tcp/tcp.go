package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/jhgv/gocodes/middleware/utils"
)

const NumRepetitions int = 10000

func startTCPServer() {
	listener, _ := net.Listen("tcp", ":8081")
	//fmt.Printf("Server running on %s", listener.Addr().String())
	var msgFromClient []byte
	for i := 0; i < NumRepetitions; i++ {
		conn, _ := listener.Accept()
		jsonDecoder := json.NewDecoder(conn)
		err := jsonDecoder.Decode(&msgFromClient)
		// log.Printf("Message from client : {%s}", msgFromClient)
		utils.CheckError(err)
		msgToClient := []byte(utils.ProcessedMessage(string(msgFromClient)))
		// log.Printf("Message to client : {%s}", msgToClient)
		jsonCoder := json.NewEncoder(conn)
		jsonCoder.Encode(msgToClient)
		conn.Close()
	}
}

func startTCPClient() {

	service := "127.0.0.1:8081"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	utils.CheckError(err)

	// Xlsx file operations
	xlsBuilder := utils.XlsxBuilder{}
	fileName := fmt.Sprintf("tcp-%d.xlsx", NumRepetitions)
	xlsBuilder.SetFileName(fileName)
	xlsBuilder.CreateHeader()
	averageFormula := fmt.Sprintf("AVERAGE(A%d:A%d)", xlsBuilder.GetRowNum()+1, NumRepetitions+2)
	xlsBuilder.SetupAverageFormula(averageFormula)
	start := time.Now()
	for i := 0; i < NumRepetitions; i++ {
		conn, _ := net.DialTCP("tcp", nil, tcpAddr)
		jsonDecoder := json.NewDecoder(conn)
		jsonCoder := json.NewEncoder(conn)
		var messageFromServer []byte
		messageToServer := []byte(utils.GenerateRandomText(200))
		// log.Printf("Message to server : {%s}", messageToServer)
		startReq := time.Now()

		_ = jsonCoder.Encode(messageToServer)
		//utils.CheckError(err)

		_ = jsonDecoder.Decode(&messageFromServer)
		elapsedReq := time.Since(startReq)
		xlsBuilder.AddRowData(elapsedReq.Seconds() * 1000.0)
		// log.Printf("Message from server : {%s}", messageFromServer)
		time.Sleep(time.Millisecond * 10)
		//utils.CheckError(err)
		conn.Close()
	}

	elapsed := time.Since(start)
	xlsBuilder.SetTotalTime(elapsed.Seconds() * 1000)
	xlsBuilder.GenerateFile()

}

func main() {
	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		startTCPServer()
	} else {
		startTCPClient()
	}
}
