package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *httprouter.Router

var handlerTestString = []struct {
	path, body string
}{
	{"/", "Welcome!"},
	{"/fibonaccisequence/10", "[0 1 1 2 3 5 8 13 21 34 55]"},
	{"/fibonaccisequence/abc", "Invaid request parameter:  abc, shall be an integer"},
	{"/fibonaccisequence/-10", "Get fibonacci sequence failed: Invalid fibonacci number!" +
		" Negative number is not allowed."},
	{"/fibonaccinumber/10", "55"},
	{"/fibonaccinumber/abc", "Invaid request parameter:  abc, shall be an integer"},
	{"/fibonaccinumber/-10", "Get fibonacci number failed: Invalid fibonacci number! " +
		"Negative number is not allowed."},
}

func TestHandlers(t *testing.T) {
	for _, v := range handlerTestString {
		req, err := http.NewRequest("GET", v.path, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("\"%v\" handler returned wrong status code: %v, it shall be %v",
				v.path, status, http.StatusOK)
		}

		// Check the response body.
		expected := v.body
		if rr.Body.String() != expected {
			t.Errorf("\"%v\" handler returned unexpected body: %v, it shall be %v",
				v.path, rr.Body.String(), expected)
		}
	}
}

func init() {
	router = initRouter()
}
