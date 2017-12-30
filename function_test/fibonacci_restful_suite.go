package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var handlerTestString = []struct {
	path, body string
}{
	{"/", "Welcome!"},
	{"/fibonaccisequence/10", "[0 1 1 2 3 5 8 13 21 34]"},
	{"/fibonaccisequence/abc", "Invaid request parameter:  abc, shall be an integer"},
	{"/fibonaccisequence/-10", "Get fibonacci sequence failed: Invalid fibonacci number!" +
		" Negative number is not allowed."},
	{"/fibonaccinumber/10", "55"},
	{"/fibonaccinumber/abc", "Invaid request parameter:  abc, shall be an integer"},
	{"/fibonaccinumber/-10", "Get fibonacci number failed: Invalid fibonacci number! " +
		"Negative number is not allowed."},
}

func TestGetOKPages() {
	for _, v := range handlerTestString {
		fmt.Printf("Testing GET \"%v\" page ... ", v.path)
		resp, err := http.Get("http://127.0.0.1:1234" + v.path)
		if err != nil {
			fmt.Println("failed: can not get the response")
			continue
		}

		// Check the status code.
		if status := resp.Status; status != "200 OK" {
			fmt.Printf("failed: got wrong status code: %v, it shall be %v\n",
				status, "200 OK")
			continue
		}

		// Check the response body.
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("failed: can not open the response body")
			continue
		}
		expected := v.body
		if string(body) != expected {
			fmt.Printf("failed: handler returned unexpected body: %v, it shall be %v\n",
				string(body), expected)
			continue
		}
		fmt.Println("passed!")
		continue
	}
	return
}

func TestPostPages() {
	for _, v := range handlerTestString {
		fmt.Printf("Testing POST \"%v\" page ... ", v.path)
		resp, err := http.Post("http://127.0.0.1:1234"+v.path, "", nil)
		if err != nil {
			fmt.Println("failed: can not get the response")
			continue
		}

		// Check the status code
		if status := resp.Status; status != "405 Method Not Allowed" {
			fmt.Printf("failed: got wrong status code: %v, it shall be %v\n",
				status, "405 Method Not Allowed")
			continue
		}

		fmt.Println("passed!")
		continue
	}
	return
}

func main() {
	TestGetOKPages()
	TestPostPages()
}
