package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

/*
	By default channels are unbuffered. Every write to an open, unbuffered channel causes
 	the writing goroutine to pause until another goroutine reads from the same channel.
 	Likewise, a read from an open, unbuffered channel causes the reading goroutine to pause
 	until another goroutine writes to the same channel. This means you cannot write to or read
 	from an unbuffered channel without at least two concurrently running goroutines.
*/
func main() {
	nums := make(chan int)
	wg.Add(2)

	fmt.Println("Starting Goroutines...")
	go consume(nums)
	go produce("A", nums)
	go produce("B", nums)

	fmt.Println("Waiting for all Goroutines...")
	wg.Wait()

	fmt.Println("Closing Channel...")
	close(nums)

	fmt.Println("Finished")
}

func consume(in chan int) {
	for num := range in {
		fmt.Println("Consumed: ", num)
	}
}

func produce(id string, in chan int) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		fmt.Printf("%s Producing -> %d \n", id, i)
		in <- i
		fmt.Printf("%s Produced -> %d \n", id, i)
	}
}
