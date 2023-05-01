package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------- Understanding Maps -----------")
	maps()
}

func maps() {

	// Uninitialized Map with nil value
	var m1 map[string]int
	fmt.Println(m1 == nil)

	// Initialized empty Map
	m2 := map[string][]string{}
	fmt.Println(m2 == nil)
	fmt.Println(m2)

	// Initialized  Map
	m3 := map[string][]string{
		"cena": []string{"a", "b"},
		"john": []string{"x", "y"},
	}
	fmt.Println(m3)
	fmt.Println(m3["cena"])

	// value check in map
	_, ok := m3["non-key"]
	if !ok {
		fmt.Println("Key is not present")
	}

	// Delete key from map
	delete(m3, "cena")
	fmt.Println(m3)
}
