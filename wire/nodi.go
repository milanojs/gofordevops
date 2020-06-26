package main

import (
	"bytes"
	"fmt"
)

type Logger struct{}

func (logger *Logger) Log(message string) {
	fmt.Println(message)
}

type HttpClient struct {
	logger *Logger
}

func (client *HttpClient) Get(url string) string {
	client.logger.Log("Getting " + url)

	// make an HTTP request
	return "my response from " + url
}

func NewHttpClient() *HttpClient {
	logger := &Logger{}
	return &HttpClient{logger}
}

type ConcatService struct {
	logger *Logger
	client *HttpClient
}

func (service *ConcatService) GetAll(urls ...string) string {
	service.logger.Log("Running GetAll")

	var result bytes.Buffer

	for _, url := range urls {
		result.WriteString(service.client.Get(url))
	}

	return result.String()
}

func NewConcatService() *ConcatService {
	logger := &Logger{}
	client := NewHttpClient()

	return &ConcatService{logger, client}
}

func main() {
	service := NewConcatService()

	result := service.GetAll(
		"http://example.com",
		"https://drewolson.org",
	)

	fmt.Println(result)
}