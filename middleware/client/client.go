package main

import (
	"github.com/jhgv/gocodes/middleware/common/naming"
	"github.com/jhgv/gocodes/middleware/core/proxy"
	"github.com/jhgv/gocodes/middleware/utils/constants"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	host     = "localhost"
	namingServerPort = 9995
	objectID = 9999
	objectName = "FileConverter"
)

func main() {
	var fileConverterProxy proxy.FileConverterProxy
	namingService := naming.NewNamingProxy(host, namingServerPort)

	absPath, _ := filepath.Abs("./middleware/client/64MB.txt")
	//file, err := ioutil.ReadFile(absPath)
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	totalTime := 0.0
	for i := 0; i < constants.NumRepetitions; i++ {
		//fmt.Printf("%d out of %d\n",i, constants.NumRepetitions)
		startReq := time.Now()
		fileConverterProxy, _ = namingService.Lookup(objectName)
		fileConverterProxy.ConvertFile(file)

		//fmt.Printf("File: \n %s", convertedFile)

		//fmt.Printf("Converted file:\n\n %s", convertedFile)

		elapsedReq := time.Since(startReq)
		totalTime = totalTime + (elapsedReq.Seconds() * 1000.0)
		time.Sleep(10 * time.Millisecond)
	}
	file.Close()
	log.Printf("Average request time: %f ms", totalTime/float64(constants.NumRepetitions))
}
