package main

import "os"

func main() {
	arg := os.Args[1]

	if len(os.Args) != 2 {
		printLn("This program only takes one input as an argument")
		return
	}
	printLn(arg)
}
