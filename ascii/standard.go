package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Standard(file []byte) {
	var line1, line2, line3, line4, line5, line6, line7, line8 string
	content := strings.Split(string(file), "\n")

	// store users input
	printable := os.Args[1]
	var a int
	var b int
	// split the users input where there is \n to exclude \n and get the other values as slices of string
	arr := strings.Split(printable, "\\n")
	for _, v := range arr {
		if len(v) > 0 {
			for _, val := range v {
				a = (int(val)-32)*9 + 1
				b = a + 7
				for x := a; x <= b; x++ {
					if x == a {
						line1 += content[x]
					} else if x == a+1 {
						line2 += content[x]
					} else if x == a+2 {
						line3 += content[x]
					} else if x == a+3 {
						line4 += content[x]
					} else if x == a+4 {
						line5 += content[x]
					} else if x == a+5 {
						line6 += content[x]
					} else if x == a+6 {
						line7 += content[x]
					} else if x == a+7 {
						line8 += content[x]
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
