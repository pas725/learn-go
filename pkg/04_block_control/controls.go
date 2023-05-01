package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------- Understanding IF, For, Labels -----------")
	ifBlock()

	forLoop()

	labels()
}

func ifBlock() {

	// Simple if block
	val := true
	if ok := val; ok {
		fmt.Println("All Ok!")
	} else if !ok {
		fmt.Println("Not Ok!")
	} else {
		fmt.Println("Ahhh!")
	}
}

func forLoop() {

	// Simple for loop (Complete)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Condition only for (like while)
	i := 1
	for i < 10 {
		fmt.Println("Conditional for", i)
		i = i * 2
	}

	// Infinite for
	i = 1
	for {
		fmt.Println("Infinite")
		i = i * 2
		if i > 10 {
			break
		}
	}

	// For range, (With index and value)
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}

// Labels can only be used with for loops
func labels() {
	i := 1
outer:
	for i < 10 {
		fmt.Println(i)
		i = i + 2
		continue outer
	}
}
