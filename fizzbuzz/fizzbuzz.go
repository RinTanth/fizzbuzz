package fizzbuzz

import "fmt"

func fb(n uint) (output string) {

	divider := []uint{15, 3, 5}
	textOutput := []string{"FizzBuzz", "Fizz", "Buzz"}

	for i, v := range divider {
		if n%v == 0 {
			return textOutput[i]
		}
	}

	return fmt.Sprintf("%d", n)
}

func FB(n uint) (output []string) {
	for i := uint(1); i <= n; i++ {
		output = append(output, fb(i))
	}
	fmt.Println(output)
	return
}
