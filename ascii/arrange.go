package ascii

func Arrange(words []string) string {
	word := ""
	for i, v := range words {
		if i != 0 {
			word += " " + string(v)
		} else {
			word += string(v)
		}
	}
	return word
}
