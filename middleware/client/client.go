package main

import (
	"log"
	"time"

	"github.com/jhgv/gocodes/middleware/core/proxy"
	"github.com/jhgv/gocodes/middleware/utils/constants"
)

const (
	host       = "localhost"
	port       = 7777
	objectName = "TextHelper"
	objectID   = 9999
)

func main() {
	textHelperProxy := &proxy.TextHelperProxy{}
	textHelperProxy.SetID(objectID)
	textHelperProxy.SetHost(host)
	textHelperProxy.SetPort(port)
	totalTime := 0.0
	for i := 0; i < constants.NumRepetitions; i++ {
		startReq := time.Now()
		textHelperProxy.UpperText("hello world")
		elapsedReq := time.Since(startReq)
		totalTime = totalTime + (elapsedReq.Seconds() * 1000.0)
		time.Sleep(10 * time.Millisecond)
	}
	log.Printf("Average request time: %f ms", totalTime/float64(constants.NumRepetitions))
}
