package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")

	fmt.Printf("Listenning at localhost:8000...\n")

	listener, _ := net.ListenTCP("tcp", addr)

	for {
		conn, _ := listener.AcceptTCP()
		go handleConn(conn)
	}

}

func handleConn(conn *net.TCPConn) {
	buffer := make([]byte, 1)

	for {
		fmt.Printf("\n[Client]: ")
		for {
			_, e := conn.Read(buffer)
			if e != nil || buffer[0] == 10 {
				break
			} else {
				fmt.Printf("%s", buffer)
			}
		}
	}

}
