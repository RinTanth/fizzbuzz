package fizzbuzz

import "strconv"

func FizzBuzz(n int) (output []string) {

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			output = append(output, "Fizz")
			continue
		}
		output = append(output, strconv.Itoa(i))
	}

	return
}
