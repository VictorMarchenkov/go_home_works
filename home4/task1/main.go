// Used idea from Mihalis Tsoukalos "Mastering Go" (2018) p. 653-658
package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

// Client structure for assigning a unique identifier to each processed request
type Client struct {
	id      int
	integer int
}

// Data structure is used for coupling the data of a Client with the actual results generated by the program
type Data struct {
	job     Client
	counter int
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
	counter = GetInstance()
)

func worker(w *sync.WaitGroup) {
	for c := range clients {
		counter.AddOne()
		output := Data{c, counter.count}
		data <- output
	}
	w.Done()
}

// makeWP generates n goroutines worker()
func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))

	nJobs := 100
	nWorkers := 5

	go create(nJobs)
	finished := make(chan interface{})

	go func() {
		for d := range data {
			fmt.Printf("Client with ID: %d\t", d.job.id)
			fmt.Printf("returns value of counter: %d\n", d.counter)
		}

		finished <- true

	}()
	makeWP(nWorkers)
	fmt.Printf(": %v\n", <-finished)
	fmt.Println(counter.count)
}
