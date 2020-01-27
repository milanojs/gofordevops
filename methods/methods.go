package main

import (
	"fmt"
	"time"
)

/*Normal Func*/
func SayHello(name string) {
	fmt.Printf("Hello %v\n", name)
}

/*return type Func*/

func getHello(name string) string {
	return fmt.Sprintf("Hello %v\n", name)
}

/*struct type return func*/
type Person struct {
	First string
	Last  string
	Dob   time.Time
}

func NewPerson(first, last string, year, month, day int) *Person {
	return &Person{
		First: first,
		Last:  last,
		Dob:   time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
	}
}

/*struct type return func*/

func (p Person) sayHelloPerson() {
	fmt.Printf("Hello, %s\n", p.First)
}

func (p Person) getAge() int {
	d := time.Since(p.Dob)
	return int(d.Hours() / 24 / 365)
}

func main() {
	SayHello("Juan")
	s := getHello("Jane")
	println(s)
	/**/

	p := NewPerson("Juan", "milano", 1986, 07, 19)
	p.sayHelloPerson()
	fmt.Println(p.getAge())
}
