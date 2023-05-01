package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------- Understanding Arrays + Slices -----------")
	arrays()
	slices()
}

func arrays() {
	fmt.Println("\n--- ARRAYS ---")
	// Array with capacity 5, intitializes all with 0
	var arr1 [5]int
	fmt.Println(arr1)

	// Array with capacity 5, intitializes with given values, sets others to 0
	var arr2 = [5]int{4, 5, 6}
	fmt.Println(arr2)

	// Array with capacity 5, intitializes with given index, sets others to 0
	var arr3 = [5]int{0: 4, 1: 5, 6}
	fmt.Println(arr3)

	// Array with inffered capacity
	var arr4 = [...]int{9, 8, 7, 6, 5}
	fmt.Println(arr4)

	// Array equality
	fmt.Println(arr2 == arr3)

	prettyPrint(&arr4)
}

func slices() {
	fmt.Println("\n--- SLICES ---")
	// Slices, backed by Array
	var sl1 = []int{4, 5, 6, 7, 8, 9, 10}
	prettyPrintSlice(&sl1)

	var sl2 = sl1[2:5]
	prettyPrintSlice(&sl2)

	// Uninitialized slice
	var sl3 []int
	fmt.Println(sl3 == nil)
	fmt.Println(len(sl3))

	sl3 = append(sl3, 9, 10, 11)
	prettyPrintSlice(&sl3)

	var tempSl = []int{1, 2, 3}
	sl3 = append(sl3, tempSl...)
	sl3[4] = 77
	sl3 = append(sl3, 99)
	prettyPrintSlice(&sl3)

	// Capacity defines total consecutive locations available, before new slice is made to add new data
	fmt.Println("Slice capacity: ", cap(sl3))

	// Create a slice with initial length, all values assigned to 0
	var sl4 = make([]int, 6)
	sl4[1] = 7
	prettyPrintSlice(&sl4)

	// Create a slice with initial length + capacity, all values assigned to 0
	var sl5 = make([]int, 5, 10)
	sl5[1] = 8
	prettyPrintSlice(&sl5)

	var emptySl = []int{}
	prettyPrintSlice(&emptySl)
}

// Print arrays
func prettyPrint(add *[5]int) {
	fmt.Println("Address:", &add, " Value:", add)
}

// Print slices
func prettyPrintSlice(add *[]int) {
	fmt.Println("Address:", &add, " Value:", add)
}
