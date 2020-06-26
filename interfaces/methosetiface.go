package main

import "fmt"

type Person struct {
	name string
	age  int
}

type NameSetter interface {
	SetName(v string)
}

func main() {
	fmt.Println("Implementing Interfaces")

	var bob = Person{"Bob Jones", 35}
	var jane = Person{"jane Jones", 35}

	fmt.Println(bob)
	fmt.Println(jane)
	jane.SetName("Mary Jane")
	fmt.Println(jane)

	pPeter := new(Person)
	pPeter.SetName("Mrs. Peters")
	fmt.Println(pPeter)
}
func (p *Person) SetName(v string) {
	p.name = v
}

func (p Person) String() string {
	return fmt.Sprintf("Person (name: %v, Age: %v)", p.name, p.age)
}
