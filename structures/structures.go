package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
	Phone Phone
}

type Phone struct {
	AreaCode string
	Preffix  string
	Suffix   string
}

func main() {
	p := Person{
		First: "john",
		Last:  "Doe",
		Age:   32,
		Phone: Phone{
			AreaCode: "+56",
			Preffix:  "9",
			Suffix:   "0",
		},
	}
	//q := &Person{"Jane", "Doe", 25}

	fmt.Println(p)
	fmt.Println("New Struct")

	pt := struct {
		x int
		y int
	}{
		x: 10,
		y: 20,
	}
	fmt.Println(pt)

}
