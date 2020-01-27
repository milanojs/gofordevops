package main

import "flag"

import "fmt"

import "os"

func main() {

	name := flag.String("name", "", "The name to say Hello to")
	flag.Parse()
	if *name == "" {
		fmt.Println("Must add name to use this tool")
		flag.Usage()
		os.Exit(1)
	}
}
