package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Hello, World!\n")

	f, e := os.OpenFile("./files/demo.txt", os.O_CREATE|os.O_RDWR, 0777)

	if e != nil {
		panic(e)
	}

	for i := 0; i < 1000; i++ {
		s := "This is line " + strconv.Itoa(i+1) + "\n"

		f.WriteString(s)
	}

	f.Close()

}
