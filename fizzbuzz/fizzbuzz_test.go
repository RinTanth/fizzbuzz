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

func TestFizzBuzz3Fizz(t *testing.T) {
	output := fizzbuzz.FizzBuzz(3)

	expected := "Fizz"

	if output != expected {
		t.Errorf("Output got %s, expected %s", output, expected)
	}

}
func TestFizzBuzz4(t *testing.T) {
	output := fizzbuzz.FizzBuzz(4)

	expected := "4"

	if output != expected {
		t.Errorf("Output got %s, expected %s", output, expected)
	}

}
