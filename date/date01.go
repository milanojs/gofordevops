package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {

	const layoutISO = "2006-01-02"

	date := "2020-0d6-15"

	match, _ := regexp.MatchString(`\d{4}\-\d{1,}\-\d{1,}`, date)

	if match == true {
		t, err := time.Parse(layoutISO, date)

		if err != nil {
			fmt.Println("There is an error validate the date: ", err)
			t = time.Now().AddDate(0, 0, 1)
		}

		loc := time.UTC
		now := time.Now().In(loc)
		futureDate := t.In(loc)

		if now.After(futureDate) {
			fmt.Println(futureDate)

			fmt.Println("After---")

		} else {
			fmt.Println(futureDate)

			fmt.Println("Before---")
		}
	}
	fmt.Println("There there was a problem with the data format ")
}
