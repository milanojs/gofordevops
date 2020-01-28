package main

import "io/ioutil"

func main() {
	err := ioutil.WriteFile("output_ioutils.txt", []byte("Hello, World!"), 0644)
	if err != nil {
		panic("Unable to write files")
	}
}
