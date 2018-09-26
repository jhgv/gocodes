package handler

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/jhgv/gocodes/middleware/rpc/upperfy"
)

type RPCClientHandler struct {
	client *rpc.Client
	reply  string
}

func (ch *RPCClientHandler) SetupSocket(host string, port int) error {
	address := fmt.Sprintf("%s:%d", host, port)
	log.Printf("address: %s", address)
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		return err
	}
	ch.client = client
	return nil
}

func (ch *RPCClientHandler) Send(message []byte) error {
	text := string(message)
	args := &upperfy.Args{Text: text}
	var reply string
	log.Printf("calling remote")
	err := ch.client.Call("Textfy.UpperText", args, &reply)
	ch.reply = reply
	if err != nil {
		return err
	}
	return nil
}

func (ch *RPCClientHandler) Recieve() ([]byte, error) {
	return []byte(ch.reply), nil
}
