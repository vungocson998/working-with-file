package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Client\n")
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")

	conn, _ := net.DialTCP("tcp", nil, addr)

	sendCmd(conn)

}

func sendCmd(conn *net.TCPConn) {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Input filepath: ")
		buffer, err := reader.ReadBytes('\n')

		if err != nil {
			log.Println("-> Input error :", err)
		} else {
			filePath := string(buffer[0 : len(buffer)-1])

			if fileExists(filePath) {

				info, _ := os.Stat(filePath)

				size := strconv.FormatInt(info.Size(), 10)

				message := "SEND" + " " + filePath + " " + size + "\n"

				conn.Write([]byte(message))

			} else {
				fmt.Printf("\nFile %s does not exist (or is a directory)\n", filePath)
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
