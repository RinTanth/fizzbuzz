package fizzbuzz

import (
	"testing"
)

func TestNormalCases(t *testing.T) {
	tests := []struct {
		input uint
		want  string
	}{
		{1, "1"},
		{2, "2"},
		{4, "4"},
		{7, "7"},
		{8, "8"},
		{11, "11"},
		{13, "13"},
		{14, "14"},
	}

	for _, tt := range tests {
		output := fb(tt.input)
		if output != tt.want {
			t.Errorf("Output got %s, want %s", output, tt.want)
		}
	}
}

func TestFizzBuzz3Fizz(t *testing.T) {
	output := fb(3)

	want := "Fizz"

	if output != want {
		t.Errorf("Output got %s, want %s", output, want)
	}

}

func TestFizzBuzz5Buzz(t *testing.T) {
	output := fb(5)

	want := "Buzz"

	if output != want {
		t.Errorf("Output got %s, want %s", output, want)
	}

}

func TestFizzBuzz6Fizz(t *testing.T) {
	output := fb(6)

	want := "Fizz"

	if output != want {
		t.Errorf("Output got %s, want %s", output, want)
	}

}

func TestFizzBuzz9Fizz(t *testing.T) {
	output := fb(6)

	want := "Fizz"

	if output != want {
		t.Errorf("Output got %s, want %s", output, want)
	}

}

func TestFizzBuzz10Buzz(t *testing.T) {
	output := fb(10)

	want := "Buzz"

	if output != want {
		t.Errorf("Output got %s, want %s", output, want)
	}

}

func TestFizzBuzz12Fizz(t *testing.T) {
	output := fb(12)

	want := "Fizz"

	if output != want {
		t.Errorf("Output got %s, want %s", output, want)
	}

}

func TestFizzBuzz15Fizz(t *testing.T) {
	output := fb(15)

	want := "FizzBuzz"

	if output != want {
		t.Errorf("Output got %s, want %s", output, want)
	}

}
