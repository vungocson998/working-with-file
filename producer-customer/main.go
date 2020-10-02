package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type product int

type producer struct {
	id int
	product
}

type customer struct {
	id int
	product
}

func main() {
	producersNumber := flag.Int("p", 1, "number of producers")
	customersNumber := flag.Int("c", 1, "number of customers")
	bufferSize := flag.Int("b", 1, "size of buffer")
	flag.Parse()

	producersList := initProducers(*producersNumber)
	customersList := initCustomers(*customersNumber)
	buffer := make(chan product, *bufferSize)

	run(producersList, customersList, buffer)

	for {
		select {
		default:
			time.Sleep(time.Millisecond * 500)
			// fmt.Printf("Main\n")
		}
	}

}

func initProducers(number int) []producer {
	producersList := make([]producer, number)
	var id int = 0

	for ; id < number; id++ {
		producersList[id] = producer{id, product(rand.Intn(10000))}
	}
	return producersList
}

func initCustomers(number int) []customer {
	customersList := make([]customer, number)
	var id int = 0

	for ; id < number; id++ {
		customersList[id] = customer{id, 0}
	}
	return customersList
}

func produce(p *producer, buffer chan product, wg *sync.WaitGroup) {
	defer wg.Done()
	buffer <- p.product
	fmt.Printf("Producer %d produced %d\n", p.id, p.product)
	p.product = 0
}

func buy(c *customer, buffer chan product, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Customer %d is waiting\n", c.id)
	c.product = <-buffer
	fmt.Printf("Customer %d bought %d\n", c.id, c.product)
}

func run(producers []producer, customers []customer, buffer chan product) {
	var wg sync.WaitGroup

	wg.Add(len(customers))
	for j := range customers {
		go buy(&customers[j], buffer, &wg)
	}
	// wg.Wait()

	wg.Add(len(producers))
	for i := range producers {
		go produce(&producers[i], buffer, &wg)
	}

}
