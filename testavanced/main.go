package main

import "fmt"

// Calculate returns x + 2.
func Add(x, y int) (result int) {
	result = x + y
	return result
}

func SayHello() {
	fmt.Println("Hello World")
}

func main() {
	SayHello()
	fmt.Printf("The add Value is: %v", Add(2, 4))
}
