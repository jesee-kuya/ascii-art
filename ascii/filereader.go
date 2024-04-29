package ascii

import (
	"errors"
	"os"
	"strings"
)

// Read the banner files and store them as slices of string
func Reader(filename string, sepp string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error reading banner file")
	}

	content := strings.Split(string(file), sepp)
	if len(content) != 856 {
		return nil, errors.New("the banner file content is not correct")
	}
	return content, nil
}
