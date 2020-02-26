package main

import "fmt"

func main() {
	dataChannel := make(chan string)
	dataChannel <- "Some Sample Data"
	fmt.Println(<-dataChannel)
}
