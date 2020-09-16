package main

import (
	"fmt"
	"io"
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
	buffer := make([]byte, 10)

	for {
		c, e := conn.Read(buffer)
		if e != nil {
			if e == io.EOF {
				fmt.Printf("Client %d disconnected!!\n", clientIndex)
				break
			} else {
				break
			}
		}

		if c > 0 && buffer[c-1] == 10 {
			fmt.Printf("%s", buffer[0:c-1])
			fmt.Printf("\t[From client %d]\n", clientIndex)
		} else {
			fmt.Printf("%s", buffer[0:c])
		}
	}

}
