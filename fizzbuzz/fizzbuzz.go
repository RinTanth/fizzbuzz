package fizzbuzz

import (
	"fmt"
)

func FizzBuzz(n int) (output []string) {

	for i := 1; i <= n; i++ {
		if i%15 != 0 {
			if i%3 == 0 {
				output = append(output, "Fizz")
				continue
			}
			if i%5 == 0 {
				output = append(output, "Buzz")
				continue
			}
			output = append(output, fmt.Sprintf("%d", i))
			continue
		}
		output = append(output, "FizzBuzz")
		continue
	}

	fmt.Println(output)
	return
}
