package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	_ "time"
)

var wg sync.WaitGroup
var wg_consumer sync.WaitGroup
var counter int32

/*
	These channels buffer a limited number of writes without blocking. If the buffer fills before there
	are any reads from the channel, a subsequent write to the channel pauses the writing goroutine until
	the channel is read. Just as writing to a channel with a full buffer blocks, reading from a channel
	with an empty buffer also blocks.
*/
func main() {
	nums := make(chan int, 20)
	wg.Add(2)

	fmt.Println("Starting Goroutines...")

	go produce("A", nums)
	go produce("B", nums)

	wg_consumer.Add(1)
	go consume(nums)

	fmt.Println("Waiting for all Goroutines...")
	fmt.Println("Channel length:", len(nums))
	fmt.Println("Channel capacity:", cap(nums))
	wg.Wait()

	fmt.Println("Closing Channel...")
	close(nums)

	wg_consumer.Wait()

	fmt.Println("Finished. Read: ", counter)
}

func consume(in chan int) {
	defer wg_consumer.Done()
	for num := range in {
		fmt.Println("Consumed: ", num)
		atomic.AddInt32(&counter, 1)
		// time.Sleep(50 * time.Millisecond)
	}
}

func produce(id string, in chan int) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		// fmt.Printf("%s Producing -> %d \n", id, i)
		in <- i
	}
}
