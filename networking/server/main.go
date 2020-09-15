package main

import (
	"fmt"
	"net"
)

func main() {
	var clientIndex int = 0
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")

	fmt.Printf("Listenning at localhost:8000...\n")

	listener, _ := net.ListenTCP("tcp", addr)

	for {
		conn, _ := listener.AcceptTCP()
		clientIndex++
		fmt.Printf("\nClient %d connected\n", clientIndex)
		go handleConn(conn, clientIndex)
	}

}

func handleConn(conn *net.TCPConn, clientIndex int) {
	buffer := make([]byte, 1)

	for {
		for {
			_, e := conn.Read(buffer)
			if e != nil || buffer[0] == 10 {
				fmt.Printf("\t[From client %d]\n", clientIndex)
				break
			} else {
				fmt.Printf("%s", buffer)
			}
		}
	}

}
