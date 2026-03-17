package fizzbuzz

import "fmt"

type Moder interface {
	Mod(x, y float64) float64
}

func FizzBuzz(mod Moder, n float64) (output []string) {

	for i := 1.00; i <= n; i++ {
		if mod.Mod(i, 15) != 0 {
			fmt.Println("mod 15 is not 0")
		}
		output = append(output, "FizzBuzz")
		continue
	}
	return
}
