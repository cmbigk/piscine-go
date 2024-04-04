package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var word string
	for _, letter := range s {
		if letter != ' ' && letter != '\t' && letter != '\n' {
			word += string(letter)
		} else if letter == ' ' || letter == '\t' || letter == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
