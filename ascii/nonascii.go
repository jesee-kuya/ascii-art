package ascii

import "fmt"


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
