package handler

import (
	"fmt"
	"net/rpc"

	"github.com/jhgv/gocodes/middleware/patterns/app"
)

type RPCClientHandler struct {
	client *rpc.Client
	reply  string
}

func (ch *RPCClientHandler) SetupSocket(host string, port int) error {
	address := fmt.Sprintf("%s:%d", host, port)
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		return err
	}
	ch.client = client
	return nil
}

func (ch *RPCClientHandler) Send(message []byte) error {
	text := string(message)
	args := &app.Args{Text: text}
	var reply string
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
