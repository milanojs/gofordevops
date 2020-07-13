package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AddResults struct {
	x        int
	y        int
	expected int
}

var addResults = []AddResults{
	{2, 4, 6},
	{-1, 1, 0},
	{0, 2, 2},
	{-5, -3, -8},
}

func TestAdd(t *testing.T) {
	for _, test := range addResults {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatalf("The sum of %v + %v should be %v", test.x, test.y, test.expected)
		}
	}
}
func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal("could no open file")
	}
	if string(data) != "hello world" {
		t.Fatal("String do not match expected")
	}
}

//mocking http response
func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"status\":\"Good\"}")
	}
	req := httptest.NewRequest("GET", "http://tutorialsedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(body))
	if 200 != resp.StatusCode {
		t.Fatal("Status Code Not ok")
	}
}
