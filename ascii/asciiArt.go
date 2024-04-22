package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Restrict the number of args to just 2
	if len(os.Args) != 2 || len(os.Args[1]) == 0 {
		return
	}
	// Get the second arg
	words := os.Args[1]

	// If the word is just a \n print newline and return
	if words == "\\n" {
		fmt.Println()
		return
	}

	// Check for an empty strng and add a \n or split using \n as seperater
	wordsArr := strings.Split(words, "\\n")

	// Read the file with ascii art handle read error then split it using \n as seperator
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Read file error", err)
	}

	fileArr := strings.Split(string(file), "\n")

	// Get the runes of wordArr
	for _, val := range wordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					fmt.Print(fileArr[int(start)+i])
				}
				fmt.Println()
			}
		} else {
			fmt.Println()
		}
	}

}
