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

func TestFizzBuzz_15(t *testing.T) {

	modVal := fizzbuzz.FizzBuzz(15)
	targetCount := 1
	count := countTarget(modVal, fb)

	if count != targetCount {
		t.Errorf("Total %s returned as %d, expected %d", fb, count, targetCount)
	}
}

func countTarget(items []string, target string) int {
	count := 0
	for _, s := range items {
		if s == target {
			count++
		}
	}
	return count
}
