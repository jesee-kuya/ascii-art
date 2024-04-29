package ascii

import (
	"errors"
	"os"
	"strings"
)

// Read reads the banner file and returns a slice of strings
// If there is an error reading the file, it returns the error.

func Reader(filename string, sepp string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error reading banner file")
	}

// If the number of lines in the file is not equal to 856, it returns an error.
	content := strings.Split(string(file), sepp)
	if len(content) != 856 {
		return nil, errors.New("the banner file content is not correct")
	}
	return content, nil
}

