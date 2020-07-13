package main

import (
	"fmt"
)

// Calculate returns x + y.
func Add(x, y int) (result int) {
	result = x + y
	return fmt.Printf("El resultado es: %d", result)
}
func SayHello() {
	fmt.Println("Hello World")
}
func main() {
	SayHello()
	Add(3, 6)
}
