package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Printf("Client\n")
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")

	conn, _ := net.DialTCP("tcp", nil, addr)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Input message: ")
		buffer, err := reader.ReadBytes('\n')

		if err != nil {
			log.Println("-> Input error :", err)
			break
		} else {
			_, err := conn.Write(buffer)
			if err != nil {
				log.Println("-> Send message error :", err)
			} else {
				log.Println("-> Send message successfully")
			}
		}
	}

	fmt.Printf("-> Close\n")
}
