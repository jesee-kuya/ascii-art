package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Reader(filename string, sepp string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	content := strings.Split(string(file), sepp)
	return content
}
