package main

import (
	"testing"
)

func checkFibonacciSequence(t *testing.T, fs []uint64, n int) {

	// Check the sequence is nil or not
	if fs == nil {
		t.Error("fibonacciSequence is nil!\n")
	}

	// Check the sequence length
	if l := len(fs); l != n {
		t.Errorf("Wrong fibonacciSequence length: %v, it shall be: v%!\n", l, n)
	}

	// Check every value in the sequence
	for i := 0; i < n; i++ {
		switch i {
		case 0:
			if v := fs[i]; v != 0 {
				t.Errorf("Wrong fibonacciSequence number: %v, it shall be: v%!\n", v, 0)
			}
		case 1:
			if v := fs[i]; v != 1 {
				t.Errorf("Wrong fibonacciSequence number: %v, it shall be: v%!\n", v, 1)
			}
		default:
			expected := fs[i-2] + fs[i-1]
			if v := fs[i]; v != expected {
				t.Errorf("Wrong fibonacciSequence number: %v, it shall be: v%!\n", v, expected)
			}
		}
	}
}

func TestInitSequence(t *testing.T) {
	checkFibonacciSequence(t, fibonacciSequence, MAXNUMBER)
}

func TestParameterCheck(t *testing.T) {
	// Check negative integer
	if parameterCheck(-1, MAXNUMBER) == nil {
		t.Error("Negative number shall not be allowed!\n")
	}

	//Check big integer
	if parameterCheck(MAXNUMBER, MAXNUMBER-1) == nil {
		t.Error("Big integar shall not be supported!\n")
	}

	//Check allowed integer
	for i := 0; i <= MAXNUMBER; i++ {
		if parameterCheck(i, MAXNUMBER) != nil {
			t.Errorf("Value %v shall be allowed!\n", i)
		}
	}
}

func TestGetFobinacciSequence(t *testing.T) {
	// Check negative integer
	if _, err := getFibonacciSequence(-1); err == nil {
		t.Error("Negative number shall not be allowed!\n")
	}

	// Check big integer
	if _, err := getFibonacciSequence(MAXNUMBER+1); err == nil {
		t.Error("Big integar shall not be supported!\n")
	}

	fs, err := getFibonacciSequence(0)
	if err != nil {
		t.Errorf("Value %v shall be allowed!\n", 0)
	}
	if len(fs) != 0 {
		t.Errorf("Empty arrary shall be returned for number 0")
	}

	// Check every allowed fibonacci sequence
	for i := 1; i <= MAXNUMBER; i++ {
		fs, err := getFibonacciSequence(i)
		if err != nil {
			t.Errorf("Value %v shall be allowed!\n", i)
		}
		checkFibonacciSequence(t, fs, i)
	}
}

func TestGetFobinacciNumber(t *testing.T) {
	// Check negative integer
	if _, err := getFibonacciNumber(-1); err == nil {
		t.Error("Negative number shall not be allowed!\n")
	}

	// Check big integer
	if _, err := getFibonacciNumber(MAXNUMBER); err == nil {
		t.Error("Big integar shall not be supported!\n")
	}

	// Check every fibonacci number
	// Construct a new fibonacci sequence for verify
	fs := make([]uint64, MAXNUMBER)
	for i := 0; i < MAXNUMBER; i++ {
		n, err := getFibonacciNumber(i)
		if err != nil {
			t.Errorf("Value %v shall be allowed!\n", i)
		}
		fs[i] = n
	}
	checkFibonacciSequence(t, fs, MAXNUMBER)
}
