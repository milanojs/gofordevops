package main

import "fmt"

func main() {
	s := "My String"
	sptr := &s
	fmt.Println(s)
	fmt.Println(sptr)  // Print memory address
	fmt.Println(*sptr) // Print pointer value via dereference
	fmt.Println("Nueva Linea")
	sptr1 := new(string)
	fmt.Println(*sptr1)

	fmt.Println("Nueva Linea 2")
	var sptr2 *string
	sptr2 = &s
	fmt.Println(*sptr2)
}
