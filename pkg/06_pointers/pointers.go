package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("----------- Understanding functions -----------")

	pointerDemo()

	loadJSON()
}

// Will not update, as it is working on the copy of address
func failedUpdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func pointerDemo() {
	x := 10
	failedUpdate(&x)
	fmt.Println(x) // prints 10
	update(&x)
	fmt.Println(x) // prints 20
}

func loadJSON() {
	f := Person{}
	err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f)

	fmt.Println(f, err)
}

// Person object
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
