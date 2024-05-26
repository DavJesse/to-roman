package main

func romanNumeros(num int) string {
	var count int
	var result string

	if num/1000 > 0 {
		count = num / 1000
		num /= 1000

		for count > 0 {
			result += "M"
			count--
		}
	}

	if num/100 > 0 {
		count = num / 100

		if count == 9 {
			result += "CM"
			num -= 900
			count = 0
		}

		if count > 5 && count < 9 {
			result += "D"
			num -= 500
			count -= 5
		}

		for count > 0 {
			if count == 4 {
				num -= 400
				result += "CD"
				count = 0
				break
			}

			num -= 100
			result += "C"
			count--
		}
	}

	if num/10 > 0 {
		count /= 10

		if count == 9 {
			result += "XC"
			num -= 90
			count = 0
		}

		if count > 5 && count < 9 {
			result += "L"
			num -= 50
		}

		for count > 0 {
			if count == 4 {
				num -= 40
				result += "XL"
				count = 0
			}
			result += "X"
			num -= 10
			count--
		}
	}

	if num/1 > 0 {
		count = num / 1

		if count == 9 {
			num -= 9
			result += "IX"
			count = 0
		}

		if count > 5 {
			count -= 5
			num -= 5
			result += "V"
		}

		for count > 0 {
			if count == 4 {
				num -= 4
				result += "IV"
				count = 0
				break
			}

			num -= 1
			result += "I"
			count--
		}
	}

	return result
}
