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
	}

	// If the cword is just a \n print newline and return
	if words == "\\n" {
		fmt.Println()
		return
	} else if len(words) == 0 {
		return
	}

	// Check for an empty strng and add a \n or split using \n as seperater
	wordsArr := strings.Split(words, "\\n")

	// Read the file with ascii art handle read error then split it using \n as seperator

	args := os.Args[1]
	if args != "-t" && args != "-s" && len(os.Args) == 2 {
		content := ascii.Reader("standard.txt", "\n")
		ascii.Ascii(content, wordsArr)
	} else if args == "-s" {
		content := ascii.Reader("shadow.txt", "\n")
		ascii.Ascii(content, wordsArr)
	} else if args == "-t" {
		content := ascii.Reader("thinkertoy.txt", "\r\n")
		ascii.Ascii(content, wordsArr)
	}
}
