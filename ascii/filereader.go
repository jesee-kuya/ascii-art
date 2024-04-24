package ascii

import (
	"fmt"
	"os"
	"strings"
)
 // Read the banner files and store them as slices of string
func Reader(filename string, sepp string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	content := strings.Split(string(file), sepp)
	return content
}
