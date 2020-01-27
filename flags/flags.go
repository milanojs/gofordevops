package main

import "flag"

import "fmt"

import "os"

import "time"

/*Variables global scope*/
var name, lastname string
var wait time.Duration
var debug bool

func main() {
	if name == "" || lastname == "" {
		fmt.Println("Must add name/lastname to use this tool")
		flag.Usage()
		os.Exit(1)
	}
	if debug {
		fmt.Printf("Going to wait for %v\n\n", wait)
	}
	time.Sleep(wait)
	fmt.Printf("Hello, %s %s\n", name, lastname)
}
func init() {
	/*Creating with a pointer*/
	//name := flag.String("name", "", "The name to say Hello to")
	/*Creating with a variable*/
	flag.StringVar(&name, "name", "", "The Last Name to say Hello to")
	flag.StringVar(&lastname, "lastname", "", "The Last Name to say Hello to")
	flag.BoolVar(&debug, "debug", false, "Turn on Debugging output")
	defaultWait, err := time.ParseDuration("5s")
	if err != nil {
		panic("There was an err")
	}
	flag.DurationVar(&wait, "wait-time", defaultWait, "time to wait before exiting")
	flag.Parse()
}
