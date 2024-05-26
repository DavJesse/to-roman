package main

import "os"

func main() {
	//Capture input to be converted
	arg := os.Args[1]

	//Exit program when user inputs excess arguments
	if len(os.Args) != 2 {
		printLn("This program only takes one input as an argument")
		return
	}

	//Convert input string to manipulable integer
	num, err := atoi(arg)

	//Exit function when input string contains non-numeric characters
	if err != "" {
		printLn(err)
		return
	}

	//Exit function when user inputs '0' or exceeds the limit of 3999
	if num >= 4000 || num == 0 {
		if num >= 4000 {
			printLn("Connot convert '" + arg + "' to Roman numerals")
			printLn("This program can only take numbers less than 4000")
		} else {
			printLn("Connot convert '" + arg + "' to Roman numerals")
			printLn("This program can only take numbers greater than '0'")
		}

		return
	}

	//Convert the integer contained in argument to roman numerals
	roman := romanNumeros(num)

	//Print result of roman conversion
	printLn(roman)

}
