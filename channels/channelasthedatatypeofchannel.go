package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func greeter(cc chan chan string) {
	c := make(chan string)
	cc <- c
}
func main() {
	fmt.Println("Main() started")

	cc := make(chan chan string)

	go greeter(cc)

	c := <-cc

	go greet(c)

	c <- "Jhon"

	fmt.Println("main() stopped")

}