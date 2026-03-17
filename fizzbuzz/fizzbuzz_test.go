package fizzbuzz_test

import (
	"fizzbuzz/fizzbuzz"
	"slices"
	"testing"
)

type mockMod struct {
	modValue float64
}

func (m *mockMod) Mod(x, y float64) float64 {
	return m.modValue
}

func TestFizzBuzz(t *testing.T) {
	mock := &mockMod{
		modValue: 0,
	}

	modVal := fizzbuzz.FizzBuzz(mock, 1)

	if !slices.Equal(modVal, []string{"FizzBuzz"}) {
		t.Errorf("ExecQuery() returned %s, expected FizzBuzz", modVal)
	}

}
