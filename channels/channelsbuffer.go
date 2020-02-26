package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v", len(c), cap(c))
	fmt.Println()
}
