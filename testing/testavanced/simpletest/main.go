package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Add func to sum 2 values
func Add(value1 int, value2 int) int {
	return value1 + value2
}

//Subtract
func Subtract(value1 int, value2 int) int {
	return value1 - value2

}

//RootEndpoint
func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe("12345", router))
}
