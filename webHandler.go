// This module is to register fibonacci web handler to the http router
// The interface shall be like:
// 	"127.0.0.1:1234/fibonaccisequence/<n>" to get the first "n" fibonacci numbers
//	"127.0.0.1:1234/fibonaccinumber/<n>" to get the "nth" fibonacci number

// The opensource package "httprouter" is used as the http request dispatcher

package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"strconv"
)


// Welcome Page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!")
}

// Handler to get the fibonacci sequence
func FibonacciSequenceHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the paramter "number" which shall be an integer
	num_string := ps.ByName("number")
	num, err := strconv.Atoi(num_string)
	if err != nil {
		fmt.Fprintf(w, "Invaid request parameter:  %v, shall be an integer", num_string)
		return
	}

	fibonacciSequence, error := getFibonacciSequence(num)
	if error != nil {
		fmt.Fprintf(w, "Get fibonacci sequence failed: %v", error)
		return
	}
	fmt.Fprint(w, fibonacciSequence)
}

// Handler to get the fibonacci number
func FibonacciNumberHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the paramter "number" which shall be an integer
	num_string := ps.ByName("number")
	num, err := strconv.Atoi(num_string)
	if err != nil {
		fmt.Fprintf(w, "Invaid request parameter:  %v, shall be an integer", num_string)
		return
	}
	fibonacciValue, error := getFibonacciNumber(num)
	if error != nil {
		fmt.Fprintf(w, "Get fibonacci number failed: %v", error)
		return
	}
	fmt.Fprint(w, fibonacciValue)
}

// Init http router and bind the handlers to it
func initRouter() *httprouter.Router {
	router := httprouter.New()
	if (router == nil) {
		log.Fatal("Can not create http router!")
	}
	router.GET("/", Index)
	router.GET("/fibonaccisequence/:number", FibonacciSequenceHandler)
	router.GET("/fibonaccinumber/:number", FibonacciNumberHandler)
	return router
}

func main (){
	router :=  initRouter()
	log.Fatal(http.ListenAndServe(":1234", router))
}
