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

func TestFizzBuzz(t *testing.T) {
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
