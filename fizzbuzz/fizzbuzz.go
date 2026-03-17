package fizzbuzz

import (
	"fmt"
	"math"
)

func FizzBuzz(n float64) (output []string) {

	for i := 1.00; i <= n; i++ {
		if math.Mod(i, 15) != 0 {
			output = append(output, fmt.Sprintf("%.0f", i))
			continue
		}
		output = append(output, "FizzBuzz")
		continue
	}
	return
}
