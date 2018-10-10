package main

import (
	"flag"
	"net"
	"os"
	"strings"
	"time"

	"github.com/jhgv/gocodes/middleware/utils"
)

// Number of repetitions
const NumRepetitions int = 10000

func startUDPServer() {
	service := ":8081"
	udpAddr, _ := net.ResolveUDPAddr("udp", service)
	conn, _ := net.ListenUDP("udp", udpAddr)
	defer conn.Close()

	for i := 0; i < NumRepetitions; i++ {

		var msgFromClient string
		buf := make([]byte, 1024)
		n, addr, _ := conn.ReadFromUDP(buf)
		msgFromClient = string(buf[0:n])

		msgToClient := strings.ToUpper(string(msgFromClient))
		conn.WriteTo([]byte(msgToClient), addr)
	}
}

func startUDPClient() {
	service := ":8081"
	// Xlsx file operations
	xlsBuilder := utils.XlsxBuilder{}
	xlsBuilder.SetBasicMetricsFile("udp", NumRepetitions)

	startTotal := time.Now()
	for i := 0; i < NumRepetitions; i++ {
		udpAddr, _ := net.ResolveUDPAddr("udp", service)
		conn, _ := net.DialUDP("udp", nil, udpAddr)

		messageToServer := []byte(utils.GenerateRandomText(100))
		startReq := time.Now()
		conn.Write(messageToServer)

		buf := make([]byte, 1024)
		conn.Read(buf)
		elapsedReq := time.Since(startReq)
		// messageFromServer := buf[:n]
		// log.Printf("Message from server: %s", messageFromServer)
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
