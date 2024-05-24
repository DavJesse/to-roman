package main

func atoi(str string) int {
	var result int
	var digit int

	for _, r := range str {
		digit = int(r - '0')
		result = result * 10 + digit
	}
	return result
} 