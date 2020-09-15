package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Printf("Client\n")
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")

	conn, _ := net.DialTCP("tcp", nil, addr)

	buffer := []byte("Hello")

	conn.Write(buffer)

	fmt.Printf("-> Close\n")
}
