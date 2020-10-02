package main

import (
	"fmt"
	"sync"
)

func main() {
	var prods, cus, channelSize int

	fmt.Printf("Number of producers: ")
	fmt.Scanf("%d\n", &prods)

	fmt.Printf("Number of customers: ")
	fmt.Scanf("%d\n", &cus)

	fmt.Printf("Channel size: ")
	fmt.Scanf("%d\n", &channelSize)

	fmt.Printf("We now have %d producers, %d customers and channel size %d\n", prods, cus, channelSize)

	run := true
	var myOption int
	var numProd, numCon int
	var channelIsEmpty bool = true
	var channelIsFull bool = false
	var channelMutex sync.Mutex

	// This is SHARED BUFFER
	myChannel := make(chan int, channelSize)

	for run == true {
		// Create a terminal interface
		fmt.Println("1. Produce a number")
		fmt.Println("2. Consume a number")
		fmt.Println("3. Exit")
		fmt.Println("Chose:")
		fmt.Scanf("%d\n", &myOption)
		switch myOption {
		case 1:
			if channelIsFull == false {
				fmt.Printf("You want to produce number: ")
				fmt.Scanf("%d\n", &numProd)
				go produce(numProd, myChannel, &channelIsFull, &channelIsEmpty, channelMutex)
			} else {
				fmt.Printf("Channel is full now!!\n")
				printChannel(myChannel)
			}
			break
		case 2:
			if channelIsEmpty == false {
				consume(&numCon, myChannel, &channelIsFull, &channelIsEmpty, channelMutex)
				fmt.Printf("You consumed number: %d\n", numCon)
			} else {
				fmt.Printf("Channel is empty now!!\n")
				printChannel(myChannel)
			}
			break
		case 3:
			fmt.Println("Good bye!!")
			run = false
			break
		}
	}
}

// Producer problems: detect if buffer is empty to produce product, notify if buffer is full after produce a product
func produce(numProd int, myChannel chan int, channelIsFull *bool, channelIsEmpty *bool, channelMutex sync.Mutex) {
	myChannel <- numProd
	channelMutex.Lock()
	defer channelMutex.Unlock()
	*channelIsEmpty = false
	if len(myChannel) == cap(myChannel) {
		*channelIsFull = true
	}
	printChannel(myChannel)
}

// Customer problems: detect if buffer is not empty to use, notify if buffer is empty after use product
func consume(numCon *int, myChannel chan int, channelIsFull *bool, channelIsEmpty *bool, channelMutex sync.Mutex) {
	x := <-myChannel
	channelMutex.Lock()
	defer channelMutex.Unlock()
	*numCon = x
	*channelIsFull = false
	if len(myChannel) == 0 {
		*channelIsEmpty = true
	}
	printChannel(myChannel)
}

func printChannel(myChannel chan int) {
	fmt.Printf("Channel has %d of %d elements", len(myChannel), cap(myChannel))
	fmt.Printf("\n==================\n\n\n")

}
