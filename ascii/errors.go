package ascii

import (
	"fmt"
	"os"
)


// NoError checks for errors in the input arguments
// and prints the correct usage if there is an error
func NoError(words []string) bool {
	// check if there is no input passed and flag without input
	// return false if the user didn't pass any input
	if len(words) <= 1 {
		fmt.Println("Error: usage; go run . <input>")
		return false
	}
	// check for incorrect flag usage
	// if the user passed a flag without input
	if (os.Args[1] == "-t" || os.Args[1] == "-s") && len(words) <= 2 {
		fmt.Println("Error: usage; go run . <flag> <input>")
		return false
	}
	// check for empty inputs
	// return false if the user passed an empty input
	if os.Args[1] == "" && len(os.Args) == 2 {
		return false
	} else if os.Args[1] == "-t" && os.Args[2] == "" && len(os.Args) == 3 {
		return false
	} else if os.Args[1] == "-s" && os.Args[2] == "" && len(os.Args) == 3 {
		return false
	}
	// check for new line flag if the user passed a flag with input as \n
	// print a new line and return false
	if os.Args[1] == "\\n" && len(os.Args) == 2 {
		fmt.Println()
		return false
	} else if os.Args[1] == "-t" && os.Args[2] == "\\n" && len(os.Args) == 3 {
		fmt.Println()
		return false
	} else if os.Args[1] == "-s" && os.Args[2] == "\\n" && len(os.Args) == 3 {
		fmt.Println()
		return false
	}
	// return true if there are no errors
	return true
}
