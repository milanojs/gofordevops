package main

import "fmt"

func greet(c chan string){
	fmt.Println("Hello " + <-c + " from Channel!")
} 
func main(){
	fmt.Println("main() started")
	c := make(chan string)
	go greet(c)
	c<- "Juan"
	fmt.Println("main() stopped")
}