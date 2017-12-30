// This module is to register fibonacci web handler to the http router
// The interface shall be like:
// 	"127.0.0.1:1234/fibonaccisequence/<n>" to get the first "n" fibonacci numbers
//	"127.0.0.1:1234/fibonaccinumber/<n>" to get the "nth" fibonacci number

// The opensource package "httprouter" is used as the http request dispatcher
// The opensource package "logrus" is used as logger

package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"net/http"
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
		log.Warnf("Invaid request parameter received:  %v", num_string)
		fmt.Fprintf(w, "Invaid request parameter:  %v, shall be an integer", num_string)
		return
	}

	log.WithField("Number", num).Info("Try to get fibonacci sequence")
	fibonacciSequence, err := getFibonacciSequence(num)
	if err != nil {
		log.WithField("Number", num).Warn("Get fibonacci sequence failed")
		fmt.Fprintf(w, "Get fibonacci sequence failed: %v", err)
		return
	}

	log.WithField("Number", num).Info("Get fibonacci sequence success")
	fmt.Fprint(w, fibonacciSequence)
}

// Handler to get the fibonacci number
func FibonacciNumberHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get the paramter "number" which shall be an integer
	num_string := ps.ByName("number")
	num, err := strconv.Atoi(num_string)
	if err != nil {
		log.Warnf("Invaid request parameter received:  %v", num_string)
		fmt.Fprintf(w, "Invaid request parameter:  %v, shall be an integer", num_string)
		return
	}

	log.WithField("Number", num).Info("Try to get fibonacci number")
	fibonacciValue, err := getFibonacciNumber(num)
	if err != nil {
		log.WithField("Number", num).Warn("Get fibonacci number failed")
		fmt.Fprintf(w, "Get fibonacci number failed: %v", err)
		return
	}

	log.WithField("Number", num).Info("Get fibonacci number success")
	fmt.Fprint(w, fibonacciValue)
}

// Init http router and bind the handlers to it
func initRouter() *httprouter.Router {
	router := httprouter.New()
	if router == nil {
		log.Fatal("Can not create http router!")
	}

	router.GET("/", Index)
	router.GET("/fibonaccisequence/:number", FibonacciSequenceHandler)
	router.GET("/fibonaccinumber/:number", FibonacciNumberHandler)
	return router
}

func main() {
	router := initRouter()
	log.Fatal(http.ListenAndServe(":1234", router))

}

func init() {
	//Todo: move the log to a file instead of stderr
	//log.SetOutput()
	log.SetLevel(log.WarnLevel)
}
