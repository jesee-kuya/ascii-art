package main

import (
	"fmt"
	"os"

	ascii "ascii/ascii"
)

func main() {

// It reads the command line arguments and performs different actions based on the flags provided.

// It checks if the words contain non-ASCII characters and prints an error if they do.
// It prints the ASCII art based on the banner file and the arranged words.
	words := os.Args
	if !ascii.NoError(words) {
		return
	}

	// sort flags with banner files
	if words[1] == "-t" {
		// It reads the banner files and stores them as slices of strings.
		content, error := ascii.Reader("thinkertoy.txt", "\r\n")
		if error != nil {
			fmt.Println(error)
			return
		}
		word := ascii.Arrange(words[2:])
		wordsArr := ascii.Slice(word)
		if !ascii.CheckAscii(wordsArr) {
			return
		}
		// It arranges the words and slices them into a 2D array.
		// It prints the ASCII art based on the banner file and the arranged words.
		ascii.Ascii(content, wordsArr)
	} else if words[1] == "-s" {
		content, error := ascii.Reader("shadow.txt", "\n")
		if error != nil {
			fmt.Println(error)
			return
		}
		word := ascii.Arrange(words[2:])
		wordsArr := ascii.Slice(word)
		if !ascii.CheckAscii(wordsArr) {
			return
		}
		ascii.Ascii(content, wordsArr)
	} else {
		content, error := ascii.Reader("standard.txt", "\n")
		if error != nil {
			fmt.Println(error)
			return
		}
		word := ascii.Arrange(words[1:])
		wordsArr := ascii.Slice(word)
		if !ascii.CheckAscii(wordsArr) {
			return
		}
		ascii.Ascii(content, wordsArr)
	}
}
