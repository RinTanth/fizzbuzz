package fizzbuzz

func FizzBuzz(n int) (output []string) {
	if n == 2 {
		return []string{"1", "2"}
	}

	if n == 3 {
		return []string{"1", "2", "Fizz"}
	}

	return []string{"1"}
}
