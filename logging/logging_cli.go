package main

import "flag"

import "log"

import "os"

import "fmt"

var name string

func main() {
	prefix := fmt.Sprintf("%s: ", os.Args[0])
	info, err := os.Create("info.log")
	if err != nil {
		log.Fatal("Failed to create log file")
	}
	infoLog := log.New(info, prefix, log.LstdFlags)
	errorLog := log.New(os.Stderr, prefix, log.LstdFlags)

	if name == "" {
		flag.Usage()
		errorLog.Println("No name supplied")
		os.Exit(1)
	}
	infoLog.Println("Program Started")
	infoLog.Printf("Hello, %s\n", name)
	infoLog.Println("Program Finished")

}
func init() {
	flag.StringVar(&name, "name", "", "The name to say hello to")
	flag.Parse()

}
