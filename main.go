package main

import (
	"fmt"
	"os"
	"strings"

	ascii "ascii/ascii"
)

func main() {
	var words string
	if len(os.Args) == 2 {
		words = os.Args[1]
	} else if len(os.Args) == 3 {
		words = os.Args[2]
	} else {
		fmt.Println("Wrong input, check README.md")
	}

	// If the word is just a \n print newline and return
	if words == "\\n" {
		fmt.Println()
		return
	} else if len(words) == 0 {
		return
	}

	// Check for an empty strng and add a \n or split using \n as seperater
	wordsArr := strings.Split(words, "\\n")

	// Read the file with ascii art handle read error then split it using \n as seperator

	if os.Args[1] != "-t" && os.Args[1] != "-s" && len(os.Args) == 2 {
		content := ascii.Reader("standard.txt", "\n")
		ascii.Ascii(content, wordsArr)
	} else if os.Args[1] == "-s" && len(os.Args) == 3 {
		content := ascii.Reader("shadow.txt", "\n")
		ascii.Ascii(content, wordsArr)
	} else if os.Args[1] == "-t" && len(os.Args) == 3 {
		content := ascii.Reader("thinkertoy.txt", "\r\n")
		ascii.Ascii(content, wordsArr)
	} else {
		fmt.Println("Wrong input, check README.md")
	}
}
