package main

import "fmt"

type BasePerson struct {
	First string
	Last  string
}
type Employee struct {
	BasePerson
	Salary int
	Boss   *Manager
}
type Manager struct {
	Employee
}
type Person interface {
	GetName() string
}

func SayHello(p Person) {
	fmt.Printf("Hello, %s\n", p.GetName())
}
func (e Employee) GetName() string {
	return e.First
}
func (m Manager) GetName() string {
	return m.First
}

func main() {
	m := &Manager{
		Employee{
			BasePerson: BasePerson{
				First: "Jhon",
				Last:  "Doe",
			},
			Salary: 40000,
			Boss:   nil,
		},
	}
	e := &Employee{
		BasePerson: BasePerson{
			First: "Steve",
			Last:  "Jones",
		},
		Salary: 30000,
		Boss:   m,
	}
	fmt.Println(m.First)
	fmt.Println(e.First)
	/*Implement person Interface*/
	SayHello(m)
	SayHello(e)
	/* Implement Slices with interfaces */
	people := []Person{Person(m), Person(e)}
	fmt.Println("Implement slices to interfaces")
	for _, person := range people {
		SayHello(person)
	}

}
