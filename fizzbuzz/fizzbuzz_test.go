package fizzbuzz_test

import (
	"fizzbuzz/fizzbuzz"
	"testing"
)

func TestFizzBuzz1(t *testing.T) {
	output := fizzbuzz.FizzBuzz(1)

	expected := "1"

	if output != expected {
		t.Errorf("Output got %s, expected %s", output, expected)
	}

}

func TestFizzBuzz2(t *testing.T) {
	output := fizzbuzz.FizzBuzz(2)

	expected := "2"

	if output != expected {
		t.Errorf("Output got %s, expected %s", output, expected)
	}

}
