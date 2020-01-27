package main

import (
	"fmt"
)

func main() {
	a := 11

	if a < 0 {
		fmt.Println("Your Value is negative")
	} else if a < 10 {
		fmt.Println("Your Value is single digit!")

	} else {
		fmt.Println("Your Value has multiple digits")
	}

}
