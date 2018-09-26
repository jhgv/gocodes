package handler

import (
	"encoding/json"
	"fmt"
	"net"
)

type TCPClientHandler struct {
	conn net.Conn
}

func (ch *TCPClientHandler) SetupSocket(host string, port int) error {
	address := fmt.Sprintf("%s:%d", host, port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	ch.conn = conn
	return nil
}

func (ch *TCPClientHandler) Send(message []byte) error {
	messageToServer := []byte(message)
	jsonCoder := json.NewEncoder(ch.conn)
	err := jsonCoder.Encode(messageToServer)
	if err != nil {
		return err
	}
	return nil
}

func (ch *TCPClientHandler) Recieve() ([]byte, error) {
	defer ch.conn.Close()
	var response string
	jsonDecoder := json.NewDecoder(ch.conn)
	err := jsonDecoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return []byte(response), nil
}
