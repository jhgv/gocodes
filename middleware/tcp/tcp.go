package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/jhgv/gocodes/middleware/models"
	"github.com/jhgv/gocodes/middleware/utils"
)

const NumRepetitions int = 5000

func startTCPServer() {
	listener, _ := net.Listen("tcp", ":8081")
	//fmt.Printf("Server running on %s", listener.Addr().String())
	var msgFromClient models.Request
	for i := 0; i < NumRepetitions; i++ {
		conn, _ := listener.Accept()
		jsonDecoder := json.NewDecoder(conn)
		err := jsonDecoder.Decode(&msgFromClient)
		utils.CheckError(err)
		msgToClient := models.Response{Message: utils.ProcessedMessage(msgFromClient.Message)}
		rep, _ := json.Marshal(msgToClient)
		conn.Write(rep)
		conn.Close()
	}
}

func startTCPClient() {

	service := "127.0.0.1:8081"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	utils.CheckError(err)
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	jsonDecoder := json.NewDecoder(conn)
	var messageFromServer models.Response

	// Xlsx file operations
	xlsBuilder := utils.XlsxBuilder{}
	fileName := fmt.Sprintf("tcp-%d.xlsx", NumRepetitions)
	xlsBuilder.SetFileName(fileName)
	xlsBuilder.CreateHeader()
	averageFormula := fmt.Sprintf("AVERAGE(A%d:A%d)", xlsBuilder.GetRowNum()+1, NumRepetitions+2)
	xlsBuilder.SetupAverageFormula(averageFormula)

	start := time.Now()
	for i := 0; i < NumRepetitions; i++ {
		messageToServer := models.Request{Message: utils.GenerateRandomText(200)}
		startReq := time.Now()
		jsonCoder := json.NewEncoder(conn)
		_ = jsonCoder.Encode(messageToServer)
		//utils.CheckError(err)

		_ = jsonDecoder.Decode(&messageFromServer)
		elapsedReq := time.Since(startReq)
		xlsBuilder.AddRowData(elapsedReq.Seconds() * 1000.0)
		time.Sleep(time.Millisecond * 10)
		//utils.CheckError(err)
	}
	elapsed := time.Since(start)
	xlsBuilder.SetTotalTime(elapsed.Seconds() * 1000)
	xlsBuilder.GenerateFile()
	conn.Close()
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
