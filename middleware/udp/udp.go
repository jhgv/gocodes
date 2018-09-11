package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"encoding/json"
	"github.com/jhgv/gocodes/middleware/utils"
	"github.com/jhgv/gocodes/middleware/models"
	"flag"
	"strings"
)

// Number of repetitions
const NumRepetitions int = 1000

func startUDPServer() {
	service := ":8888"
	udpAddr, _ := net.ResolveUDPAddr("udp", service)
	conn, _ := net.ListenUDP("udp", udpAddr)
	defer conn.Close()

	for i := 0; i < NumRepetitions; i++ {
		var msgFromClient models.Request
		var buf [1048]byte
		n, addr, err := conn.ReadFromUDP(buf[0:])
		utils.CheckError(err)
		json.Unmarshal(buf[0:n], &msgFromClient)
		msgToClient := models.Response{Message: utils.ProcessedMessage(msgFromClient.Message)}
		if err != nil {
			return
		}
		resp, _ := json.Marshal(msgToClient)
		conn.WriteTo(resp, addr)

		//jsonDecoder.Decode(&msgFromClient)
		////utils.CheckError(err)
		//msgToClient := models.Response{Message: utils.ProcessedMessage(msgFromClient.Message)}
		//resp, _ := json.Marshal(msgToClient)
		//conn.WriteTo(resp, udpAddr)
	}
}

func startUDPClient() {
	service := ":8888"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	utils.CheckError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	jsonDecoder := json.NewDecoder(conn)
	utils.CheckError(err)
	defer conn.Close()

	// Xlsx file operations
	xlsBuilder := utils.XlsxBuilder{}
	fileName := fmt.Sprintf("udp-%d.xlsx", NumRepetitions)
	xlsBuilder.SetFileName(fileName)
	xlsBuilder.CreateHeader()
	averageFormula := fmt.Sprintf("AVERAGE(A%d:A%d)", xlsBuilder.GetRowNum() + 1, NumRepetitions + 2)
	xlsBuilder.SetupAverageFormula(averageFormula)

	startTotal := time.Now()
	for i := 0; i < NumRepetitions; i++ {
		var messageFromServer models.Response
		messageToServer := models.Request{Message: utils.GenerateRandomText(400)}
		startReq := time.Now()

		jsonCoder := json.NewEncoder(conn)
		_ = jsonCoder.Encode(messageToServer)
		//utils.CheckError(err)
		_ = jsonDecoder.Decode(&messageFromServer)
		//utils.CheckError(err)

		elapsedReq := time.Since(startReq)
		//log.Printf("Message from server: %s", messageFromServer)
		xlsBuilder.AddRowData(elapsedReq.Seconds() * 1000)
		time.Sleep(time.Millisecond * 10)
	}

	elapsedTotal := time.Since(startTotal)
	xlsBuilder.SetTotalTime(elapsedTotal.Seconds() * 1000)
	xlsBuilder.GenerateFile()

	os.Exit(0)
}

func main() {
	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		startUDPServer()
	} else {
		startUDPClient()
	}

}
