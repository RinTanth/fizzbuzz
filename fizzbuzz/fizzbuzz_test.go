package fizzbuzz_test

import (
	"fizzbuzz/fizzbuzz"
	"testing"
)

var (
	fb = "FizzBuzz"
	f  = "Fizz"
	b  = "Buzz"
)

func TestFizzBuzz1(t *testing.T) {
	output := fizzbuzz.FizzBuzz(1)

	expected := []string{"1"}

	if len(output) != len(expected) {
		t.Errorf("length mismatch: got %d, want %d", len(output), len(expected))
		return
	}

	for i, v := range output {
		if v != expected[i] {
			t.Errorf("at index %d: got %q, want %q", i, v, expected[i])
		}
	}
}

func TestFizzBuzz2(t *testing.T) {
	output := fizzbuzz.FizzBuzz(2)

	expected := []string{"1", "2"}

	if len(output) != len(expected) {
		t.Errorf("length mismatch: got %d, want %d", len(output), len(expected))
		return
	}

	for i, v := range output {
		if v != expected[i] {
			t.Errorf("at index %d: got %q, want %q", i, v, expected[i])
		}
	}
}

func TestFizzBuzz3(t *testing.T) {
	output := fizzbuzz.FizzBuzz(3)

	expected := []string{"1", "2", "Fizz"}

	if len(output) != len(expected) {
		t.Errorf("length mismatch: got %d, want %d", len(output), len(expected))
		return
	}

	for i, v := range output {
		if v != expected[i] {
			t.Errorf("at index %d: got %q, want %q", i, v, expected[i])
		}
	}
}

func TestFizzBuzz4(t *testing.T) {
	output := fizzbuzz.FizzBuzz(4)

	expected := []string{"1", "2", "Fizz", "4"}

	if len(output) != len(expected) {
		t.Errorf("length mismatch: got %d, want %d", len(output), len(expected))
		return
	}

	for i, v := range output {
		if v != expected[i] {
			t.Errorf("at index %d: got %q, want %q", i, v, expected[i])
		}
	}
}

func TestFizzBuzz5(t *testing.T) {
	output := fizzbuzz.FizzBuzz(5)

	expected := []string{"1", "2", "Fizz", "4", "Buzz"}

	if len(output) != len(expected) {
		t.Errorf("length mismatch: got %d, want %d", len(output), len(expected))
		return
	}

	for i, v := range output {
		if v != expected[i] {
			t.Errorf("at index %d: got %q, want %q", i, v, expected[i])
		}
	}
}

func TestFizzBuzz6(t *testing.T) {
	output := fizzbuzz.FizzBuzz(6)

	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz"}

	if len(output) != len(expected) {
		t.Errorf("length mismatch: got %d, want %d", len(output), len(expected))
		return
	}

	for i, v := range output {
		if v != expected[i] {
			t.Errorf("at index %d: got %q, want %q", i, v, expected[i])
		}
	}
}
