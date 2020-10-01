package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	EOL_CHAR_1 = byte('\n')
	EOL_CHAR_2 = byte('\r')
)

func main() {
	filePath := flag.String("file", "../files/receive.txt", "path to the file")
	n := flag.Int("n", 1, "number of last lines you want to print out")
	flag.Parse()
	tailer(*n, *filePath)
}

func tailer(n int, filePath string) {
	fmt.Printf("Last %d lines of %s:\n\n", n, filePath)

	var num int
	var i int64
	var prevIsBreak bool = true // Used to detect \r\n or \n\r cases
	buffer := make([]byte, 1)

	f, _ := os.OpenFile(filePath, os.O_RDONLY, 0777)
	defer f.Close()
	f.Seek(0, 2)

	num = 0 // Used to track if detected line number = expected lines number
	i = 0

	// After this loop file pointer will move to expectation position (in front of n last lines)
	for num < n {
		f.Seek(-i, 2)
		f.Read(buffer)
		if buffer[0] != EOL_CHAR_1 && buffer[0] != EOL_CHAR_2 {
			prevIsBreak = false
			i++
			continue
		}
		if !prevIsBreak {
			num++
			prevIsBreak = true
		}
		i++
	}
	// Print n lines
	for {
		_, e := f.Read(buffer)
		if e == io.EOF {
			break
		}
		fmt.Printf("%s", buffer)
	}
	fmt.Printf("\n")
}
