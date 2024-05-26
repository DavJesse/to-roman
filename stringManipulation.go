package main

// Converts string to integer and capture invalid characters
func atoi(str string) (int, string) {
	var result int
	var digit int
	var errStr string

	//Range over string, finding apppropriate and inappropriate characters
	for i, r := range str {
		if i == 0 && r == '+' || r == '-' {

			//Identify negative numbers
			if r == '-' {
				errStr += "This program only covert positive integers to Roman numbers"
			}
			continue
		}

		//Identify decimal point numbers
		if r == '.' {
			errStr = errStr + "\n" + "This program only converts whole numbers to their Roman representation"
		}

		//Convert the runes to digits and eventually to an integer
		digit = int(r - '0')
		result = result*10 + digit
	}

	return result, errStr
}
