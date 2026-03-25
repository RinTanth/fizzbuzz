package fizzbuzz

import "strconv"

func fb(n int) (output string) {

	if n%15 == 0 {
		return "FizzBuzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	if n%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(n)
}

func FB(n int) (output []string) {
	for i := 1; i <= n; i++ {
		output = append(output, fb(i))
	}
	return
}
