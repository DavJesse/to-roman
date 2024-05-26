package main

import "os"

func main() {
	arg := os.Args[1]

	if len(os.Args) != 2 {
		printLn("This program only takes one input as an argument")
		return
	}
	num, err := atoi(arg)

	if err != "" {
		printLn(err)
		return
	}

	if num >= 4000 || num == 0{
		if num >= 4000 {
			printLn("Connot convert '" + arg + "' to Roman numerals")
			printLn("This program has an input limit of 4000")
		} else {
			printLn("Connot convert '" + arg + "' to Roman numerals")
			printLn("This program only takes positive integers greater than '0'")
		}
		
		return
	}

	roman := romanNumeros(num)

	printLn(roman)

}
