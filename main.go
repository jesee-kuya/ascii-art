package main

import (
	"os"

	ascii "ascii/ascii"
)

func main() {
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		panic(err)
	}
	file1, err := os.ReadFile("shadow.txt")
	if err != nil {
		panic(err)
	}
	file2, err := os.ReadFile("thinkertoy.txt")
	if err != nil {
		panic(err)
	}
	args := os.Args[1]

	if args != "-t" && args != "-s" {
		ascii.Standard(file)
	} else if args == "-s" {
		ascii.Shadow(file1)
	} else if args == "-t" {
		ascii.Thinkertoy(file2)
	}
}
