package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Shadow(file []byte) {
	var line1, line2, line3, line4, line5, line6, line7, line8 string
	content := strings.Split(string(file), "\n")
	printable := os.Args[2]
	arr := strings.Split(printable, "\\n")

	for _, v := range arr {
		if len(v) > 0 {

			for _, val := range v {
				upperLimit := (int(val)-32)*9 + 1
				lowwerLimit := upperLimit + 7
				for index := upperLimit; index <= lowwerLimit; index++ {
					if index == upperLimit {
						line1 += content[index]
					} else if index == upperLimit+1 {
						line2 += content[index]
					} else if index == upperLimit+2 {
						line3 += content[index]
					} else if index == upperLimit+3 {
						line4 += content[index]
					} else if index == upperLimit+4 {
						line5 += content[index]
					} else if index == upperLimit+5 {
						line6 += content[index]
					} else if index == upperLimit+6 {
						line7 += content[index]
					} else if index == upperLimit+7 {
						line8 += content[index]
					}
				}
			}
			fmt.Println(line1)
			fmt.Println(line2)
			fmt.Println(line3)
			fmt.Println(line4)
			fmt.Println(line5)
			fmt.Println(line6)
			fmt.Println(line7)
			fmt.Println(line8)
			line1, line2, line3, line4, line5, line6, line7, line8 = "", "", "", "", "", "", "", ""
		} else {
			fmt.Println()
		}
	}
}
