package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Printf("Client\n")
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")

	conn, _ := net.DialTCP("tcp", nil, addr)

	sendCmd(conn)

}

func sendCmd(conn *net.TCPConn) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Input filepath: ")
	buffer, err := reader.ReadBytes('\n')

	if err != nil {
		log.Println("-> Input error :", err)
	} else {
		filePath := string(buffer[0 : len(buffer)-1])

		if fileExists("../../files/client/" + filePath) {

			info, _ := os.Stat("../../files/client/" + filePath)

			size := strconv.FormatInt(info.Size(), 10)

			message := "SEND" + " " + filePath + " " + size + "\n"

			conn.Write([]byte(message))

			time.Sleep(5000 * time.Millisecond)

			fileTransfer(filePath, conn)

		} else {
			fmt.Printf("\nFile %s does not exist (or is a directory)\n", filePath)
		}
	}

}

func fileTransfer(filePath string, conn *net.TCPConn) {
	fp, _ := os.OpenFile("../../files/client/"+filePath, os.O_RDONLY, 0777)

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
			// log.Println("-> Send data successfully")
		}
	}

	fp.Close()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
