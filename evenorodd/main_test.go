package main

import "testing"

func TestMain(t *testing.T) {
	if checkEvenOdd(3) != false {
		t.Error("3 should be an add number but it was said to be even")
	}
}
