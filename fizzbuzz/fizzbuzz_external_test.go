package fizzbuzz_test

import (
	"fizzbuzz/fizzbuzz"
	"reflect"
	"testing"
)

func TestBussiness5(t *testing.T) {
	got := fizzbuzz.FB(5)

	want := 5

	if len(got) != want {
		t.Errorf("got's length got %d, want %d", len(got), want)

	}
}

func TestBussiness5got(t *testing.T) {
	got := fizzbuzz.FB(5)

	want := []string{"1", "2", "Fizz", "4", "Buzz"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %#v, want %#v\n", got, want)

	}
}

func TestBusiness0(t *testing.T) {
	got := fizzbuzz.FB(0)

	want := 0

	if len(got) != want {
		t.Errorf("got's length got %d, want %d", len(got), want)
	}
}
