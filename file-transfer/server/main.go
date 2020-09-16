package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
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

	fp, _ := os.OpenFile("../../files/receive.txt", os.O_CREATE|os.O_WRONLY, 0777)

	defer fp.Close()

	for {
		c, e := conn.Read(buffer)
		if e != nil {
			if e == io.EOF {
				fp.Write(buffer[0:c])
				log.Printf("\nClient %d disconnected!!\n", clientIndex)
				break
			} else {
				break
			}
		} else {
			fp.Write(buffer[0:c])
		}
	}
}
