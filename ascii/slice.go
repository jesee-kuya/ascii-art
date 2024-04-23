package ascii

import (
	"regexp"
	"strings"
)

func Slice(word string) []string {
	if strings.Contains(word, "\\t") {
		re := regexp.MustCompile(`\\t`)
		word = re.ReplaceAllString(word, "    $0")
		word = re.ReplaceAllString(word, "")
		wordArr := strings.Split(word, "\\n")
		return wordArr
	} else if strings.Contains(word, "\\v") {
		re := regexp.MustCompile(`\\v`)
		word = re.ReplaceAllString(word, "$0\\n    ")
		word = re.ReplaceAllString(word, "")
		wordArr := strings.Split(word, "\\n")
		return wordArr
	} else {
		wordArr := strings.Split(word, "\\n")
		return wordArr
	}
}
