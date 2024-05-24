package main

import "os"

func printLn(s string) {
	os.Stdout.WriteString(s + "\n")
}
