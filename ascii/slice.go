package ascii

import (
	"regexp"
	"strings"
)

func Slice(word string) []string {
	re := regexp.MustCompile(`\\t`)
	word = re.ReplaceAllString(word, "    $0")
	word = re.ReplaceAllString(word, "")

	re1 := regexp.MustCompile(`\\v|\\f`)
	word = re1.ReplaceAllString(word, "$0\\n    ")
	word = re1.ReplaceAllString(word, "")

	re2 := regexp.MustCompile(`(.*?)(\\r)`)
	word = re2.ReplaceAllString(word, "")
	wordArr := strings.Split(word, "\\n")
	return wordArr
}
