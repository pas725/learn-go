// Note: main package is the entrypoint to run
package main

import (
	"fmt"
	_ "time" // Way to keep unused imports
)

// Multiple variables at the same time
var (
	x             int8   = 12
	name, address string = "Cena", "US"
	defaultInt    int    // Will be assigned default zero value
)

// Ways to declare constant
const (
	MaxAge int = 50

	// Note: Constant type of struct is not allowed
	// person Person = Person{"John"}
)

// Note: Entrypoint function
func main() {
	// Note: GO is call by value
	newVar := 2.6 // Assigns type and value
	fmt.Println("----------- Understanding types & declaration -----------")

	fmt.Println("Shorthand variable: ", newVar)
}

// Person Defining struct(User define types)
type Person struct {
	name string
}
