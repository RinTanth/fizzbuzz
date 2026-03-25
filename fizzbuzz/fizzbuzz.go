package fizzbuzz

import "strconv"

func FizzBuzz(n int) (output string) {
	if n%3 == 0 {
		return "Fizz"
	}

	if n == 5 {
		return "Buzz"
	}

	return strconv.Itoa(n)
}
