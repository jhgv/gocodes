package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"
	"encoding/json"
	"github.com/jhgv/gocodes/middleware/utils"
	"github.com/jhgv/gocodes/middleware/models"
)

const NumRepetitions int = 10000

func startTCPServer() {
	listener, _ := net.Listen("tcp", ":8081")
	//fmt.Printf("Server running on %s", listener.Addr().String())
	conn, _ := listener.Accept()
	defer conn.Close()
	jsonDecoder := json.NewDecoder(conn)
	var msgFromClient models.Request

	for i := 0; i < NumRepetitions; i++ {
		err := jsonDecoder.Decode(&msgFromClient)
		utils.CheckError(err)
		msgToClient := models.Response{Message: utils.ProcessedMessage(msgFromClient.Message)}
		rep, _ := json.Marshal(msgToClient)
		conn.Write(rep)
	}
}

func startTCPClient() {

	service := "127.0.0.1:8081"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	utils.CheckError(err)
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	jsonDecoder := json.NewDecoder(conn)
	var messageFromServer models.Response

	// Xlsx file operations
	xlsBuilder := utils.XlsxBuilder{}
	fileName := fmt.Sprintf("tcp-%d.xlsx", NumRepetitions)
	xlsBuilder.SetFileName(fileName)
	xlsBuilder.CreateHeader()
	averageFormula := fmt.Sprintf("AVERAGE(A%d:A%d)", xlsBuilder.GetRowNum() + 1, NumRepetitions + 2)
	xlsBuilder.SetupAverageFormula(averageFormula)

	start := time.Now()
	for i := 0; i < NumRepetitions; i++ {
		messageToServer := models.Request{Message: utils.GenerateRandomText(100)}
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
