package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------- Understanding functions -----------")

	fmt.Println("Variadic params:", add(3, 1, 1, 1, 1))

	id, name := multiReturn()
	fmt.Println("Multiple return values:", id, name)

	testFuncAsVal()

	// Anonymous Functions
	func() {
		fmt.Println("Anonymous function.")
	}() // Call it

	deferExample()
}

// Variadic params
func add(a int, b ...int) int {
	sum := a
	for _, v := range b {
		sum += v
	}
	return sum
}

// Multiple Return Values
func multiReturn() (int, string) {
	return 5, "John"
}

// Functions Are Values
func testFuncAsVal() {
	// assign function to variable
	newAdd := add
	fmt.Println(newAdd(1, 2, 3))
}

// Defer executes in LIFO order, used to cleanup resources
func deferExample() {
	defer func() {
		fmt.Println("Closing file...")
	}()

	defer func() {
		fmt.Println("Closing DB...")
	}()

	fmt.Println("Processing file...")
	fmt.Println("Processing DB...")

}
