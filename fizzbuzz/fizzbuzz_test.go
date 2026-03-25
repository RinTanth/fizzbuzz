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

func TestFizzBuzz5Buzz(t *testing.T) {
	output := fizzbuzz.FizzBuzz(5)

	expected := "Buzz"

	if output != expected {
		t.Errorf("Output got %s, expected %s", output, expected)
	}

}

func TestFizzBuzz6Fizz(t *testing.T) {
	output := fizzbuzz.FizzBuzz(6)

	expected := "Fizz"

	if output != expected {
		t.Errorf("Output got %s, expected %s", output, expected)
	}

}

func TestFizzBuzz7(t *testing.T) {
	output := fizzbuzz.FizzBuzz(7)

	expected := "7"

	if output != expected {
		t.Errorf("Output got %s, expected %s", output, expected)
	}

}
func TestFizzBuzz8(t *testing.T) {
	output := fizzbuzz.FizzBuzz(8)

	expected := "8"

	if output != expected {
		t.Errorf("Output got %s, expected %s", output, expected)
	}

}

func TestFizzBuzz9Fizz(t *testing.T) {
	output := fizzbuzz.FizzBuzz(6)

	expected := "Fizz"

	if output != expected {
		t.Errorf("Output got %s, expected %s", output, expected)
	}

}

func TestFizzBuzz10Buzz(t *testing.T) {
	output := fizzbuzz.FizzBuzz(10)

	expected := "Buzz"

	if output != expected {
		t.Errorf("Output got %s, expected %s", output, expected)
	}

}
