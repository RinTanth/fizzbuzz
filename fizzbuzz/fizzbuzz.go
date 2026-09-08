package fizzbuzz

import "fmt"

func fb(n uint) string {
	fizz := map[bool]string{true: "Fizz", false: ""}
	buzz := map[bool]string{true: "Buzz", false: ""}
	text := fizz[n%3 == 0] + buzz[n%5 == 0]

	num := map[bool]string{true: fmt.Sprintf("%d", n), false: ""}

	return num[text == ""] + text
}

func FB(n uint) (output []string) {
	for i := uint(1); i <= n; i++ {
		output = append(output, fb(i))
	}
	fmt.Println(output)
	return
}
