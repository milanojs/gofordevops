package main

import (
	"fmt"
	"net/http"
)
type Response struct {
    Status string
    StatusCode int
    Proto string
    ProtoMajor int
    ProtoMinor int
    Header Header
    Body io.ReadCloser
    ContentLength int64
    TransferEncoding []string
    Close bool
    Uncompressed bool
    Trailer Header
    Request *Request
    TLS *tls.ConnectionState
}

func main(){
	r, _ := http.Get("http://www.example.com/index.html")
	fmt.Println(r)
}