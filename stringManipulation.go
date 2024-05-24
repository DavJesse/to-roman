package main

func atoi(str string) (int, string) {
	var result int
	var digit int
	var errStr string

	for i, r := range str {
		if i == 0 && r == '+' || r == '-' {
			if r == '-' {
				errStr += "This program only covert positive integers to Roman numbers"
			}
			continue
		}
		if r == '.' {
			errStr = errStr + "\n" + "This program only converts whole numbers to their Roman representation"
		}
		digit = int(r - '0')
		result = result * 10 + digit
	}
	return result, errStr
} 