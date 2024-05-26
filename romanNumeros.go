package main

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

func romanNumeros(num int) string {
	var result string

	for _, val := range roman {
		for num >= val.value {
			result += val.symbol
			num -= val.value
		}
	}
	return result
}
