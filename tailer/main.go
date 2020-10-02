package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	EndOfLineChar1 = byte('\n')
	EndOfLineChar2 = byte('\r')
)

func main() {
	filePath := flag.String("file", "../files/receiver.txt", "path to the file")
	n := flag.Int("n", 1, "number of last lines you want to print out")
	flag.Parse()
	tailer(*n, *filePath)
}

func tailer(expectedLines int, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		return fmt.Errorf("can not open file, check if %s is a file or directory", filePath)
	}
	defer f.Close()

	var currentLines int = 0
	var prevIsBreak bool = true
	var offsetFromBottom int64 = 0
	buffer := make([]byte, 1)

	for currentLines < expectedLines {
		f.Seek(-offsetFromBottom, 2)
		f.Read(buffer)
		if buffer[0] != EndOfLineChar1 && buffer[0] != EndOfLineChar2 {
			prevIsBreak = false
			offsetFromBottom++
			continue
		}
		if !prevIsBreak {
			currentLines++
			prevIsBreak = true
		}
		offsetFromBottom++
	}
	fmt.Printf("Last %d lines of %s:\n\n", expectedLines, filePath)
	for {
		_, e := f.Read(buffer)
		if e != nil {
			break
		}
		fmt.Printf("%s", buffer)
	}
	return nil
}
