package main

import (
	"fmt"
	"github.com/jhgv/gocodes/middleware/common/naming"
	"github.com/jhgv/gocodes/middleware/core/proxy"
	"github.com/jhgv/gocodes/middleware/utils/constants"
	"log"
	"os"
	"path/filepath"
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
 	fileConverterProxy, _ = namingService.Lookup(objectName)

	//totalTime := 0.0
	for i := 0; i < constants.NumRepetitions; i++ {
		//startReq := time.Now()
		absPath, _ := filepath.Abs("./middleware/client/100kb.txt")
		//file, err := ioutil.ReadFile(absPath)
		file, err := os.Open(absPath)
		if err != nil {
			log.Fatal(err)
		}

		convertedFile := fileConverterProxy.ConvertFile(file)

		fmt.Printf("File: \n %s", convertedFile)

		//fmt.Printf("Converted file:\n\n %s", convertedFile)

		//fileInfo, _ := convertedFile.Stat()
		//fmt.Printf("Received file: %s", fileInfo.Name())
		//textHelperProxy.UpperText("hello world")
		//elapsedReq := time.Since(startReq)
		//totalTime = totalTime + (elapsedReq.Seconds() * 1000.0)
		//time.Sleep(10 * time.Millisecond)

	}
	//log.Printf("Average request time: %f ms", totalTime/float64(constants.NumRepetitions))
}
