package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	tt := []struct {
		name   string
		value  string
		double int
		err    string
	}{
		{name: "double of two", value: "4", double: 8},
		{name: "Missing value", value: "", err: "missing value"},
		{name: "not a number", value: "x", err: "not a number"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "http://localhost:9090/double?v="+tc.value, nil)
			if err != nil {
				t.Fatalf("could not create request: %v ", err)
			}
			rec := httptest.NewRecorder()
			doubleHandler(rec, req)
			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Could not read response %v", err)
			}

			if tc.err != "" {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status bad request; got %v", res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("Expected message %q; got %q", tc.err, msg)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("Expected Status OK; got %v", res.StatusCode)
			}

			d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
			if err != nil {
				t.Fatalf("Expected an integer; got %s -", err)
			}
			if d != tc.double {
				t.Fatalf("Expected double to be %v; got %d - %T", tc.double, d, d)
			}
		})
	}
}

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()
	res, err := http.Get(fmt.Sprintf("%s/double?v=2", srv.URL))
	if err != nil {
		t.Fatalf("could not set GET request:%v", err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Could not read response %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected Status OK; got %v", res.StatusCode)
	}

	d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
	if err != nil {
		t.Fatalf("Expected an integer; got %s -", err)
	}
	if d != 4 {
		t.Fatalf("Expected double to be 4; got %d - %T", d, d)
	}
}
