package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	elapsedTime := time.Since(start)

	fmt.Println("Total Time For Execution: " + elapsedTime.String())

	time.Sleep(time.Second)
}
