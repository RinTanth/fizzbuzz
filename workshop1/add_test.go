package workshop1

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAddCaseEmpty(t *testing.T) {
	result, _ := Add("")

	assert.Equal(t, result, 0)
}

func TestAddCaseOneString(t *testing.T) {
	result, _ := Add("1")

	assert.Equal(t, result, 1)
}

func TestAddCaseTwoStrings(t *testing.T) {
	result, _ := Add("1,2")

	assert.Equal(t, result, 3)
}

func TestAddAnyAmountOfNumbers(t *testing.T) {
	result, _ := Add("1,2,3,4")

	assert.Equal(t, result, 10)
}

func TestNewlinesAsDelimiters(t *testing.T) {
	result, _ := Add("1\n2,3")

	assert.Equal(t, result, 6)
}

func TestCustomDelimiters(t *testing.T) {
	result, _ := Add("//;\n1;2")

	assert.Equal(t, result, 3)
}

func TestNewlinesAsDelimitersV2(t *testing.T) {
	result, _ := Add("1\n2,10")

	assert.Equal(t, result, 13)
}

func TestNegativeShouldError(t *testing.T) {
	_, error := Add("-1,2")

	assert.Equal(t, error, errors.New("negative numbers not allowed"))
}

func FuzzAdd(f *testing.F) {
	f.Add("1")
	f.Add("1,2")
	f.Fuzz(func(t *testing.T, s string) {
		got, err := Add(s)
		if err != nil && got == 0 {
			t.Fatalf("neg: %d", got)
		}
	})
}
