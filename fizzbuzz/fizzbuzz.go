package fizzbuzz

import (
	"math"
)

func FizzBuzz(n float64) (output []string) {

	for i := 1.00; i <= n; i++ {
		if math.Mod(i, 15) != 0 {
			if math.Mod(i, 3) == 0 {
				output = append(output, "Fizz")
				continue
			}
		}
		output = append(output, "FizzBuzz")
		continue
	}
	return
}
