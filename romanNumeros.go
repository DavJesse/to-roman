package main

// Establish a reference for important Roman numerals
var roman = []struct {
	value       int
	symbol      string
	calculation string
}{
	{1000, "M", "M"},
	{900, "CM", "(M-C)"},
	{500, "D", "D"},
	{400, "CD", "(D-C)"},
	{100, "C", "C"},
	{90, "XC", "(C-X)"},
	{50, "L", "L"},
	{40, "XL", "(L-X)"},
	{10, "X", "X"},
	{9, "IX", "(X-I)"},
	{5, "V", "V"},
	{4, "IV", "(V-I)"},
	{1, "I", "I"},
}

// Converts integers to their respective Roman numerals
func romanNumeros(num int) (string, string) {
	var result string
	var calc string

	// Range over the struct, 'roman'
	for _, val := range roman {
		// As long as 'num' is greater or equal to the values in the struct...
		for num >= val.value {

			// Add the symbol to result..
			result += val.symbol

			// Capture the method used to calculate the roman numeral
			if calc != "" {
				calc += "+" + val.calculation
			} else {
				calc += val.calculation
			}

			// Update number by subtracting the value that matched it
			num -= val.value
		}
	}

	return result, calc
}
