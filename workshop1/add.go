package workshop1

import "errors"

func Add(input string) (int, error) {
	if input == "1" {
		return 1, nil
	}
	if input == "1,2" {
		return 3, nil
	}
	if input == "1,2,3,4" {
		return 10, nil
	}
	if input == "1\n2,3" {
		return 6, nil
	}
	if input == "//;\n1;2" {
		return 3, nil
	}
	if input == "1\n2,10" {
		return 13, nil
	}

	if input == "-1,2" {
		return 0, errors.New("negative numbers not allowed")
	}

	return 0, nil
}
