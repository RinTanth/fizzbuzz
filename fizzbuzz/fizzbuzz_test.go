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

func TestFizzBuzz7(t *testing.T) {
	output := fizzbuzz.FizzBuzz(7)

	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7"}

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

func TestFizzBuzz8(t *testing.T) {
	output := fizzbuzz.FizzBuzz(8)

	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8"}

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

func TestFizzBuzz9(t *testing.T) {
	output := fizzbuzz.FizzBuzz(9)

	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz"}

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

func TestFizzBuzz10(t *testing.T) {
	output := fizzbuzz.FizzBuzz(10)

	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz"}

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

func TestFizzBuzz11(t *testing.T) {
	output := fizzbuzz.FizzBuzz(11)

	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11"}

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

func TestFizzBuzz12(t *testing.T) {
	output := fizzbuzz.FizzBuzz(12)

	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz"}

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

func TestFizzBuzz13(t *testing.T) {
	output := fizzbuzz.FizzBuzz(13)

	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13"}

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

func TestFizzBuzz14(t *testing.T) {
	output := fizzbuzz.FizzBuzz(14)

	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14"}

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
