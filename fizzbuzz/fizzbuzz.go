package fizzbuzz

import "fmt"

func fb(n uint) (output string) {

	if n%3 == 0 {
		output += "Fizz"
	}

	if n%5 == 0 {
		output += "Buzz"
	}

	switch output {
	case "":
		return fmt.Sprintf("%d", n)
	default:
		return output
	}

}

func FB(n uint) (output []string) {
	for i := uint(1); i <= n; i++ {
		output = append(output, fb(i))
	}
	return
}
