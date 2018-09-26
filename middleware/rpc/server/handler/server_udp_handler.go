package handler

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type UDPServerRequestHanlder struct {
	conn       *net.UDPConn
	listener   net.Listener
	updAddress *net.UDPAddr
}

func (ch *UDPServerRequestHanlder) SetupSocket(host string, port int) error {
	address := fmt.Sprintf(":%d", port)
	udpAddr, _ := net.ResolveUDPAddr("udp", address)
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	ch.conn = conn
	log.Printf("Server listening on port %d...\n", port)
	return nil
}

func (ch *UDPServerRequestHanlder) Send(message []byte) error {
	msgToClient := strings.ToUpper(string(message))
	_, err := ch.conn.WriteTo([]byte(msgToClient), ch.updAddress)
	if err != nil {
		return err
	}

	return nil
}

func (ch *UDPServerRequestHanlder) Recieve() ([]byte, error) {
	var msgFromClient string
	buf := make([]byte, 1024)
	n, addr, err := ch.conn.ReadFromUDP(buf)
	if err != nil {
		return nil, err
	}
	ch.updAddress = addr
	msgFromClient = string(buf[0:n])
	log.Print(msgFromClient)
	fmt.Printf("received: %s from: %s\n", msgFromClient, addr)
	return []byte(msgFromClient), nil
}
