package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	futureDate := time.Date(2020, 03, 01, 00, 00, 00, 00, time.UTC)

	if now.After(futureDate) {
		fmt.Println("true")
	} else {
		fmt.Println("False")
	}

}
