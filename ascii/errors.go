package ascii

import (
	"fmt"
	"os"
)

func NoError(words []string) bool {
	// check if there is no input passed and flag without input
	if len(words) <= 1 {
		fmt.Println("Error: usage; go run . <input>")
		return false
	} else if (os.Args[1] == "-t" || os.Args[1] == "-s") && len(words) <= 2 {
		fmt.Println("Error: usage; go run . <flag> <input>")
		return false
	}
	// check for empty inputs
	if os.Args[1] == "" && len(os.Args) == 2 {
		return false
	} else if os.Args[1] == "-t" && os.Args[2] == "" && len(os.Args) == 3 {
		return false
	} else if os.Args[1] == "-s" && os.Args[2] == "" && len(os.Args) == 3 {
		return false
	}
	// check for new line
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
	return true
}
