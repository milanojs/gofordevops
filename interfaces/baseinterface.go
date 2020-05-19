//Basic simple interface in golang

package main

import "fmt"

//Person have a name
type Person struct {
	name string
	age  int
}
type IndianRunner struct {
	name   string
	weight int
}
type WaterVale struct {
	name   string
	weight int
}
type Duck interface {
	Quack()
	Waddle()
}

func main() {
	fmt.Println("Interface type")
	var bob = Person{"Bob james", 35}
	var jane = Person{"Jane Doe", 91}

	fmt.Println(bob)
	fmt.Println(jane)

	var duck0 = IndianRunner{"Mosty", 14}
	fmt.Println(duck0)

	var duck1 = WaterVale{"Rusty", 20}
	fmt.Println(duck1)

	useDuck(duck0)

}
func (p Person) String() string {
	return fmt.Sprintf("Person (name: %v, Age: %v)", p.name, p.age)
}

func (d IndianRunner) String() string {
	return fmt.Sprintf("My Indian Runner (name: %v, weight: %v)", d.name, d.weight)
}

func (d WaterVale) String() string {
	return fmt.Sprintf("My Water Vale (name: %v, weight: %v)", d.name, d.weight)
}

func useDuck(d Duck) {
	d.Quack()
	d.Waddle()
}
func (d IndianRunner) Quack() {
	fmt.Printf("My Indian Runner is Quacking\n")
}
func (d IndianRunner) Waddle() {
	fmt.Printf("My Indian Runner is waddleing\n")
}
