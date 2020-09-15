package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")

	listener, _ := net.ListenTCP("tcp", addr)

	for {

		conn, _ := listener.AcceptTCP()

		fmt.Printf("Listenning at localhost:8000...\n")

		go handleConn(conn)
	}

}

func handleConn(conn *net.TCPConn) {
	buffer := make([]byte, 1)

	for {
		_, e := conn.Read(buffer)
		if e != nil {
			break
		} else {
			fmt.Printf("%s", buffer)
		}
	}
}
