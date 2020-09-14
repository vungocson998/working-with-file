package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

func main() {
	start := time.Now()

	numbers := readFile("../../files/input.txt")
	fmt.Printf("Array of %d numbers :\n%v\n", len(numbers), numbers)

	sort.Ints(numbers)
	fmt.Printf("Sorted array: %v\n", numbers)

	elapsed := time.Since(start)
	fmt.Printf("Time: %s\n", elapsed)
}

func readFile(filePath string) (numbers []int) {
	f, e := os.OpenFile(filePath, os.O_RDONLY, 0777)

	defer f.Close()

	if e != nil {
		panic(e)
	}

	var num int

	for {
		_, err := fmt.Fscanf(f, "%d\t", &num)

		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				return
			}
			panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))

		}
		numbers = append(numbers, num)
	}
	return
}
