package main

import "fmt"

func main() {
	m := map[string]string{}
	m["first"] = "Jhon"
	m["last"] = "Doe"
	fmt.Println(m)

	/* another way to */
	m1 := make(map[string]string)
	m1["first"] = "Jane"
	m1["last"] = "Doe"
	fmt.Println(m1)

}
