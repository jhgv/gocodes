package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/jhgv/gocodes/middleware/patterns/utils/protocols"
)

type TCPServerRequestHanlder struct {
	conn     net.Conn
	listener net.Listener
}

func (ch *TCPServerRequestHanlder) SetupSocket(host string, port int) error {
	address := fmt.Sprintf(":%d", port)
	listener, _ := net.Listen(protocols.TCP, address)
	// if err != nil {
	// 	return err
	// }
	ch.listener = listener
	log.Printf("Server listening on port %d...\n", port)
	return nil
}

func (ch *TCPServerRequestHanlder) Send(message []byte) error {
	jsonCoder := json.NewEncoder(ch.conn)
	uMessage := strings.ToUpper(string(message))
	jsonCoder.Encode(uMessage)
	// if err != nil {
	// 	return err
	// }
	ch.conn.Close()
	return nil
}

func (ch *TCPServerRequestHanlder) Recieve() ([]byte, error) {
	conn, _ := ch.listener.Accept()
	// if err != nil {
	// 	return nil, err
	// }
	ch.conn = conn
	var msgFromClient []byte
	jsonDecoder := json.NewDecoder(ch.conn)
	jsonDecoder.Decode(&msgFromClient)
	// if err != nil {
	// 	return nil, err
	// }
	return msgFromClient, nil
}
