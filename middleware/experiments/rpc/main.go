package main

import (
	"flag"
	"github.com/jhgv/gocodes/middleware/experiments/rpc/client"
	"github.com/jhgv/gocodes/middleware/experiments/rpc/server"
	"strings"
)

const (
	protocol     = "rpc"
	host         = "localhost"
	port     int = 7777
)

func main() {
	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		server.StartServer(port)
	} else {
		client.StartClient(host, port)
	}
}
