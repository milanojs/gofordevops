package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	futureDate := time.Date(2020, 03, 01, 00, 00, 00, 00, time.UTC)
	fmt.Printf("###############################################################\n")

	diff := futureDate.Sub(now)
	if futureDate == now {
		fmt.Println("es la fecha actual", int(diff.Hours()/24))
	} else {
		fmt.Println("no es la fecha actual")
	}

	daystil := int(diff.Hours() / 24)

	if daystil > 0 {
		fmt.Println("es la fecha actual %d", int(diff.Hours()/24))
	} else {
		fmt.Println("no es la fecha actual dias til", daystil)
	}

	if now.After(futureDate) {
		fmt.Println("true")
	} else {
		fmt.Println("False")
	}

}
