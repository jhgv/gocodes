package client

import (
	"fmt"
	"github.com/jhgv/gocodes/middleware/experiments/rpc/core"
	"github.com/jhgv/gocodes/middleware/experiments/utils/constants"
	"log"
	"net/rpc"
	"os"
	"path/filepath"
	"time"
)

func StartClient(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	totalTime := 0.0
	var reply []byte
	for i := 0; i < constants.NumRepetitions; i++ {
		//fmt.Printf("%d out of %d\n",i, constants.NumRepetitions)
		startReq := time.Now()
		client, err := rpc.DialHTTP("tcp", address)
		if err != nil {
			log.Fatal("Error starting connection: ", err)
		}

		absPath, _ := filepath.Abs("./middleware/client/64MB.txt")
		//file, err := ioutil.ReadFile(absPath)
		file, err := os.Open(absPath)
		if err != nil {
			log.Fatal(err)
		}

		fileInfo, _ := file.Stat()
		b := make([]byte, fileInfo.Size())
		file.Read(b)


		args := &core.Args{File: b}
		client.Call("FileConverter.ConvertFile", args, &reply)
		//fmt.Printf("File: \n %s", reply)

		elapsedReq := time.Since(startReq)
		totalTime = totalTime + (elapsedReq.Seconds() * 1000.0)
		client.Close()
		time.Sleep(10 * time.Millisecond)

		file.Close()
	}
	log.Printf("Average request time: %f ms for %d requests", totalTime/float64(constants.NumRepetitions), constants.NumRepetitions)
}
