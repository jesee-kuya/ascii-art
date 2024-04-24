package main

import (
	"os"

	ascii "ascii/ascii"
)

func main() {
	words := os.Args

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
