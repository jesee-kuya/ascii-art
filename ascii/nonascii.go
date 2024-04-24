package ascii

import "fmt"

// check for non ascii character and prints an error
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
