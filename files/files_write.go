package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		panic("Unable to create file")
	}
	defer f.Close()
	cnt, err := f.WriteString("Hello World!")
	if err != nil {
		panic("Unable to write file")
	}
	fmt.Printf("Wrote %d bytes \n", cnt)
}
