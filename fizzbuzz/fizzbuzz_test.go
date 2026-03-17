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

func TestFizzBuzz_3(t *testing.T) {

	modVal := fizzbuzz.FizzBuzz(3)
	targetCount := 1
	count := countTarget(modVal, f)

	if count != targetCount {
		t.Errorf("Total %s returned as %d, expected %d", f, count, targetCount)
	}

	if !checkLocation(modVal, f, 2) {
		t.Errorf("Location %d is %v, expected %s", 2, modVal[2], f)
	}
}

func TestFizzBuzz_5(t *testing.T) {

	modVal := fizzbuzz.FizzBuzz(5)
	targetCount := 1
	count := countTarget(modVal, f)

	if count != targetCount {
		t.Errorf("Total %s returned as %d, expected %d", b, count, targetCount)
	}

	if !checkLocation(modVal, b, 4) {
		t.Errorf("Location %d is %v, expected %s", 4, modVal[4], b)
	}
}

func TestFizzBuzz_15(t *testing.T) {

	modVal := fizzbuzz.FizzBuzz(15)

	targetWords := []string{f, b, fb}
	targetCounts := []int{4, 2, 1}
	counts := []int{countTarget(modVal, f), countTarget(modVal, b), countTarget(modVal, fb)}

	for i := 0; i < 3; i++ {
		if counts[i] != targetCounts[i] {
			t.Errorf("Total %s returned as %d, expected %d", targetWords[i], counts[i], targetCounts[i])
		}
	}

	if !checkLocation(modVal, fb, 14) {
		t.Errorf("Location %d is %v, expected %s", 14, modVal[14], fb)
	}

	if !checkMultiLocation(modVal, f, []int{2, 5, 8}) {
		t.Errorf("Fizz locations mismatched")
	}

	if !checkMultiLocation(modVal, b, []int{4, 9}) {
		t.Errorf("Buzz Locations mismatched")
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

func checkMultiLocation(items []string, target string, locations []int) bool {
	for _, l := range locations {
		if items[l] == target {
			continue
		}
		return false
	}
	return true
}
