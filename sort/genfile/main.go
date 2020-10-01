package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	filePath := flag.String("file", "../../files/input.txt", "where you want to save your file")
	num := flag.Int("n", 10000, "amount of numbers you want to generate to the file")
	flag.Parse()
	generateFile(*filePath, *num)
}

// Generate random numbers which are separated by "\t"
func generateFile(filePath string, num int) {
	f, _ := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	defer f.Close()
	var s string

	for i := 0; i < num; i++ {
		s = strconv.Itoa(rand.Intn(20000)) + "\t"
		f.WriteString(s)
	}

	fmt.Printf("Exported file %s\n", filePath)
}
