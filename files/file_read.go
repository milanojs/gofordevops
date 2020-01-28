package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("output.txt")
	if err != nil {
		panic("Unable to open file")
	}
	defer f.Close()
	buff := make([]byte, 64)
	cnt, err := f.Read(buff)
	if err != nil {
		panic("Unable to read file")
	}
	fmt.Printf("Read %d bytes \n", cnt)
	fmt.Println(string(buff[:cnt]))

}
