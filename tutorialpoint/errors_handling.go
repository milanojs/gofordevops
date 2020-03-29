package main

import "errors"
import "fmt"
import "math"

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}
func CheckErrors(value float64, err2 error) {
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(value)
	}
}

func main() {

	result, err := Sqrt(-1)
	CheckErrors(result, err)
	/*
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	*/
	result, err = Sqrt(9)
	CheckErrors(result, err)

	/*
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	*/
}
