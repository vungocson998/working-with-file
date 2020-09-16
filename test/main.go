package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("../files/demo.txt")

	fmt.Println("File is not exist!", os.IsNotExist(err))
	fmt.Println("File size : ", info.Size())
}
