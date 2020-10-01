package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func tailer(n int, filePath string) {
	fmt.Printf("Last %d lines :\n\n", n)

	var num int
	var i int64
	var prevIsBreak bool = true // Used to detect \r\n or \n\r cases
	buffer := make([]byte, 1)

	// Open and move fseek to the end of file
	f, _ := os.OpenFile(filePath, os.O_RDONLY, 0777)
	f.Seek(0, 2)

	num = 0 // Used to track if detected line number = expected lines number
	i = 0

	// After this loop file pointer will move to expectation position (in front of n last lines)
	for num < n {
		f.Seek(-i, 2)
		f.Read(buffer)

		if buffer[0] == byte(10) || buffer[0] == byte(13) {
			// Used to detect \r\n or \n\r cases
			if prevIsBreak {

			} else {
				num++
				prevIsBreak = true
			}
		} else {
			prevIsBreak = false
		}
		i++
	}

	// Print n lines
	for true {
		_, e := f.Read(buffer)
		if e == io.EOF {
			break
		} else {
			fmt.Printf("%s", buffer)
		}
	}

	fmt.Printf("\n")

}

func main() {
	n, _ := strconv.Atoi(os.Args[1])

	filePath := os.Args[2]

	tailer(n, filePath)

}
