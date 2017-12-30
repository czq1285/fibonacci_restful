// This module is to generate a static fibonacci sequence
// and provide the methods to fetch the requested sequence or number

package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

// A static variable to store the pre-generated fibonacci sequence.
var fibonacciSequence []uint64

// Maximum fibonacci number supported by uint64
// Todo: Try another type instead of unit64 to support big integer.
const MAXNUMBER = 94

// Initialize the static fibonacci number
func initFibonacciSequence() {
	log.Debug("Start to init static fibonacci sequence")
	fibonacciSequence = make([]uint64, MAXNUMBER)
	for i := 0; i < MAXNUMBER; i++ {
		switch i {
		// fibonacci[0] = 0
		case 0:
			fibonacciSequence[i] = 0
		// fibonacci[1] = 1
		case 1:
			fibonacciSequence[i] = 1
		// fibonacci[n] = fibonacci[n-2] + fibonacci[n-1]
		default:
			fibonacciSequence[i] = fibonacciSequence[i-2] + fibonacciSequence[i-1]
		}
	}
	log.WithField("Length", len(fibonacciSequence)).Debug("Init static fibonacci sequence finished")
}

// Check the parameter is valid or not
func parameterCheck(n, max int) error {
	// Check the request number, it shall not be a negative number or larger than 93.
	switch {
	case n < 0:
		log.WithField("Number", n).Warn("Negative number received")
		return errors.New("Invalid fibonacci number! Negative number is not allowed.")
	case n > max:
		log.WithField("Number", n).Warnf("Maximum number %v exceeded!", max)
		return errors.New("This fibonacci number is not supported so far. Please try smaller number.")
	}
	return nil
}

// Get the fibonacci sequence by number
func getFibonacciSequence(n int) ([]uint64, error) {
	// Check the parameter is valid or not
	if err := parameterCheck(n, MAXNUMBER); err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"Number": n,
		"Result": fibonacciSequence[:n],
	}).Debug("Get fibonacci sequence success.")
	return fibonacciSequence[:n], nil
}

// Get the fibonacci number
func getFibonacciNumber(n int) (uint64, error) {
	// Check the parameter is valid or not
	if err := parameterCheck(n, MAXNUMBER-1); err != nil {
		return 0, err
	}

	log.WithFields(log.Fields{
		"Number": n,
		"Result": fibonacciSequence[n],
	}).Debug("Get fibonacci number success.")
	return fibonacciSequence[n], nil
}

// Init the fibonacci sequence during system start
func init() {
	initFibonacciSequence()
}
