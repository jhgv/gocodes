package handler

import (
	"fmt"
	"net/rpc"

	"github.com/jhgv/gocodes/middleware/rpc/upperfy"
)

type RPCClientHandler struct {
	client *rpc.Client
}

func (ch *RPCClientHandler) SetupSocket(host string, port int) error {
	address := fmt.Sprintf("%s:%s", host, port)
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
	err := ch.client.Call("Textfy.UpperText", args, &reply)
	if err != nil {
		return err
	}
	return nil
}

func (ch *RPCClientHandler) Recieve() ([]byte, error) {
	return nil, nil
}
