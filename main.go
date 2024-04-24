package main

import (
	"fmt"
	"os"

	ascii "ascii/ascii"
)

func main() {
	words := os.Args
	// check if there is no input passed and flag without input
	if len(words) <= 1 {
		fmt.Println("Error, No arguments passed")
		return
	} else if (os.Args[1] == "-t" || os.Args[1] == "-s") && len(words) <= 2 {
		fmt.Println("Error, No arguments passed")
		return
	}
	// check for empty inputs
	if os.Args[1] == "" && len(os.Args) == 2 {
		return
	} else if os.Args[1] == "-t" && os.Args[2] == "" && len(os.Args) == 3 {
		return
	} else if os.Args[1] == "-s" && os.Args[2] == "" && len(os.Args) == 3 {
		return
	}

	// sort flags with banner files
	if words[1] == "-t" {
		content := ascii.Reader("thinkertoy.txt", "\r\n")
		word := ascii.Arrange(words[2:])
		wordsArr := ascii.Slice(word)
		if !ascii.CheckAscii(wordsArr) {
			return
		}
		ascii.Ascii(content, wordsArr)
	} else if words[1] == "-s" {
		content := ascii.Reader("shadow.txt", "\n")
		word := ascii.Arrange(words[2:])
		wordsArr := ascii.Slice(word)
		if !ascii.CheckAscii(wordsArr) {
			return
		}
		ascii.Ascii(content, wordsArr)
	} else {
		content := ascii.Reader("standard.txt", "\n")
		word := ascii.Arrange(words[1:])
		wordsArr := ascii.Slice(word)
		if !ascii.CheckAscii(wordsArr) {
			return
		}
		ascii.Ascii(content, wordsArr)
	}
}
