package main

import (
	"errors"
	"fmt"
)

func main() {
	total, err := sumOf(1, 8)
	if err != nil {
		fmt.Println("There is an error: ", err)
	} else {
		fmt.Println("La suma es", total)
	}

}
func sumOf(start, end int) (int, error) {
	var total int
	if start > end {
		return 0, errors.New("start is bigger than end")
	}
	for i := start; i <= end; i++ {
		total += i
	}
	return total, nil
}
