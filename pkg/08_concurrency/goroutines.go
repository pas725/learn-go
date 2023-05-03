package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Num of cores: ", runtime.NumCPU())

	wg.Add(2)

	// Start a goroutine
	go prime("A")
	go prime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

func prime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 15; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
