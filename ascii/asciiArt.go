package ascii

import (
	"fmt"
)

func Ascii(fileArr []string, wordsArr []string) {
	// Get the runes of wordArr and print the ascii-art
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
