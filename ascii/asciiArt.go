package ascii

import (
	"fmt"
)


// Ascii prints the ASCII art based on the banner file and the input words.
// It takes two slices as input: fileArr and wordsArr.
func Ascii(fileArr []string, wordsArr []string) {
	
	for _, val := range wordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					// Print the corresponding character of the banner file at the
					// calculated index
					fmt.Print(fileArr[int(start)+i])
				}
				// Print a newline after each line of the ASCII character of the word
				fmt.Println()
			}
		} else {
			// If val is empty, print a new line
			fmt.Println()
		}
	}
}
