package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
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
	var cmdRecv string

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
			cmdRecv = cmdRecv + string(buffer[0:c-1])
			// fmt.Printf("%s", buffer[0:c-1])
			// fmt.Printf("\t[From client %d]\n", clientIndex)
			// fmt.Printf("MESSAGE: %s\n", cmdRecv)
			cmd := strings.Split(cmdRecv, " ")
			size, _ := strconv.ParseInt(cmd[2], 10, 64)
			switch cmd[0] {
			case "SEND":
				log.Println("Handling SEND cmd")
				go handleSend(conn, clientIndex, cmd[1], size)
				cmd = nil
			}

		} else {
			cmdRecv = cmdRecv + string(buffer[0:c])
			// fmt.Printf("%s", buffer[0:c])
		}
	}

}

func handleSend(conn *net.TCPConn, clientIndex int, filePath string, size int64) {
	if fileExists("../../files/server/" + filePath) {

		log.Printf("File %s existed, file size %d!!\n", filePath, size)

	} else {
		// conn.Write([]byte("OK"))

		log.Printf("File %s is transfering!!\n", filePath)

		buffer := make([]byte, 10)

		fp, _ := os.OpenFile("../../files/server/"+filePath, os.O_CREATE|os.O_WRONLY, 0777)

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
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
