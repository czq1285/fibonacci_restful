// This module is to generate a static fibonacci sequence
// and provide the methods to fetch the requested sequence or number

package main

import (
	"errors"
)


// A static variable to store the pre-generated fibonacci sequence.
var fibonacciSequence []uint64

// Maximum fibonacci number supported by uint64
// Todo: Try another type instead of unit64 to support big integer.
const MAXNUMBER = 94

// Initialize the static fibonacci number
func initFibonacciSequence() {
	fibonacciSequence = make([]uint64, MAXNUMBER)
	for i := 0; i < MAXNUMBER; i++{
		switch i {
		// fibonacci[0] = 0
		case 0: fibonacciSequence[i] = 0
		// fibonacci[1] = 1
		case 1: fibonacciSequence[i] = 1
		// fibonacci[n] = fibonacci[n-2] + fibonacci[n-1]
		default:
			fibonacciSequence[i] = fibonacciSequence[i-2] + fibonacciSequence[i-1]
		}
	}
}

// Check the parameter is valid or not
func parameterCheck(n int) error {
	// Check the request number, it shall not be a negative number or larger than 93.
	switch {
	case n < 0:
		return errors.New("Invalid fibonacci number! Negative number is not allowed.")
	case n >= MAXNUMBER:
		return errors.New("This fibonacci number is not supported so far. Please try smaller number.")
	}
	return nil
}

// Get the fibonacci sequence by number
func getFibonacciSequence (n int) ([]uint64, error) {
	// Check the parameter is valid or not
	if err := parameterCheck(n); err != nil {
		return nil, err
	}

	return fibonacciSequence[:n+1], nil
}

// Get the fibonacci number
func getFibonacciNumber (n int) (uint64, error) {
	// Check the parameter is valid or not
	if err := parameterCheck(n); err != nil {
		return 0, err
	}

	return fibonacciSequence[n], nil
}


// Init the fibonacci sequence during system start
func init() {
	initFibonacciSequence()
}