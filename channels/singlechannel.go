package main

import "fmt"

func main() {
	dataChannel := make(chan string)
	fmt.Println(<-dataChannel)
}
