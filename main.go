package main

import (
	"fmt"
	"os"

	ascii "ascii/ascii"
)

func main() {
	words := os.Args
	if !ascii.NoError(words) {
		return
	}

	// sort flags with banner files
	if words[1] == "-t" {
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
