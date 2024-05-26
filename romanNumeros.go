package main

// Establish a reference for important Roman numerals
var roman = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// Converts integers to their respective Roman numerals
func romanNumeros(num int) string {
	var result string

	//Range over the struct, 'roman'
	for _, val := range roman {

		//As long as 'num' is greater or equal to the values in the struct...
		for num >= val.value {

			//Add the symbol to result..
			result += val.symbol

			//Update number by subtracting the value that matched it
			num -= val.value
		}
	}

	return result
}
