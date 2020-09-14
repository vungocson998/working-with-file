package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	f, _ := os.OpenFile("../../files/input.txt", os.O_CREATE|os.O_RDWR, 0777)

	var s string

	for i := 0; i < 10000; i++ {
		s = strconv.Itoa(rand.Intn(20000)) + "\t"
		f.WriteString(s)
	}

	f.Close()
	fmt.Printf("Exported file ../files/input.txt\n")
}
