package main

import "io/ioutil"

import "fmt"

func main() {
	buf, err := ioutil.ReadFile("output_ioutils.txt")
	if err != nil {
		panic("Unable to open file")
	}
	fmt.Println(string(buf))
}
