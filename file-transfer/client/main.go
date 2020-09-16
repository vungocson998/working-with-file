package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Printf("Client\n")
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")

	conn, _ := net.DialTCP("tcp", nil, addr)

	fp, _ := os.OpenFile("../../files/demo.txt", os.O_RDONLY, 0777)

	buffer := make([]byte, 10)

	for {
		c, e := fp.Read(buffer)
		if e != nil {
			if e == io.EOF {
				_, err := conn.Write(buffer[0:c])
				if err != nil {
					log.Println("-> Send data error :", err)
				} else {
					log.Println("-> Send data successfully")
				}
				log.Println("End of file")
			} else {
				log.Println("Error while reading file!")
			}
			break
		}

		_, err := conn.Write(buffer[0:c])
		if err != nil {
			log.Println("-> Send data error :", err)
		} else {
			log.Println("-> Send data successfully")
		}
	}

	fp.Close()
	fmt.Printf("-> Close\n")
}
