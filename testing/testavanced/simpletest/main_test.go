package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"
	//"github.com/tj/assert"
)

func TestAdd(t *testing.T) {
	total := Add(1, 3)
	assert.NotNil(t, total, "The total is 4")
	assert.Equal(t, 4, total, "Expecting 4")

}

func TestSubtract(t *testing.T) {
	total := Subtract(1, 3)
	assert.NotNil(t, total, "The total is -2")
	assert.Equal(t, -2, total, "Expecting 2")
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	return router
}

func TestRootEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "ok Status is Expected")
	assert.Equal(t, "Hello World", response.Body.String(), "expected Hellow world")
}
