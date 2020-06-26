package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Printer interface {
	Print()
}

func main() {
	fmt.Println("Implementing Interfaces empty")

	var bob = Person{"Bob Jones", 35}
	var jane = Person{"jane Jones", 35}

	fmt.Println(bob)
	fmt.Println(jane)

	jane.SetName("Mary Jane")
	fmt.Println(jane)

}

func (p Person) String() string {
	return fmt.Sprintf("Person (name: %v, Age: %v)", p.name, p.age)
}
