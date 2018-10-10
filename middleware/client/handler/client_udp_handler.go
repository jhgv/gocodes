package handler

import (
	"fmt"
	"net"
)

type UDPClientHandler struct {
	conn net.Conn
}

func (ch *UDPClientHandler) SetupSocket(host string, port int) error {
	address := fmt.Sprintf("%s:%d", host, port)
	udpAddr, _ := net.ResolveUDPAddr("udp", address)
	// if err != nil {
	// 	return err
	// }
	conn, _ := net.DialUDP("udp", nil, udpAddr)
	// if err != nil {
	// 	return err
	// }
	ch.conn = conn
	return nil
}

func (ch *UDPClientHandler) Send(message []byte) error {
	ch.conn.Write(message)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (ch *UDPClientHandler) Recieve() ([]byte, error) {
	defer ch.conn.Close()
	buf := make([]byte, 1024)
	n, _ := ch.conn.Read(buf)
	// if err != nil {
	// 	return nil, err
	// }
	response := buf[:n]
	// if err != nil {
	// 	return nil, err
	// }
	return response, nil
}