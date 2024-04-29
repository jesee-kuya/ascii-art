package ascii

import "fmt"


// CheckAscii checks if the input words contain any non-ASCII character
// and returns true if there are no non-ASCII characters and false otherwise.
func CheckAscii(word []string) bool {
	for _, v := range word {
		for _, val := range v {
			if val < 32 || val > 126 {
				fmt.Println("found a non-ascii character")
				return false
			}
		}
	}
	return true
}
