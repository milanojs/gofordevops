package interface

import "fmt"

func main() {
	/*Empty Interfaces*/
	var x interface{}
	x = "Hello, World!"
	/*Print empty ones*/
	s, ok := x.(string)
	if !ok {
		panic("No Its, not!")
	}
	fmt.Println(s)

	/*Using interfaces as base for switch command */
	switch x.(type) {
	case int:
		fmt.Println("Integer!")
	case string:
		fmt.Println("String!")

	}
}
