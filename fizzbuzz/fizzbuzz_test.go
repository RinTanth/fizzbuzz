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

	if !checkLocation(modVal, fb, 14) {
		t.Errorf("Location %d is %v, expected %s", 14, modVal[14], fb)
	}
}

func TestFizzBuzz_3(t *testing.T) {

	modVal := fizzbuzz.FizzBuzz(3)
	targetCount := 1
	count := countTarget(modVal, f)

	if count != targetCount {
		t.Errorf("Total %s returned as %d, expected %d", f, count, targetCount)
	}

	if !checkLocation(modVal, fb, 2) {
		t.Errorf("Location %d is %v, expected %s", 2, modVal[2], f)
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

func checkLocation(items []string, target string, location int) bool {
	return items[location] == target
}
