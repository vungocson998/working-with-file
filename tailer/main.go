package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])

	fmt.Printf("Last %d lines :\n\n", n)

	var num int
	var i int64
	var prevIsBreak bool = true
	buffer := make([]byte, 1)

	f, _ := os.OpenFile("../files/demo.txt", os.O_RDONLY, 0777)

	f.Seek(0, 2)

	num = 0
	i = 0

	for num < n {
		f.Seek(-i, 2)
		f.Read(buffer)

		if buffer[0] == byte(10) || buffer[0] == byte(13) {
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
