package fizzbuzz

import "fmt"

func fb(n uint) (output string) {

	if n%15 == 0 {
		return "FizzBuzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	if n%5 == 0 {
		return "Buzz"
	}

	return fmt.Sprintf("%d", n)
}

func FB(n uint) (output []string) {
	for i := uint(1); i <= n; i++ {
		output = append(output, fb(i))
	}
	return
}
