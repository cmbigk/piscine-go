package piscine

func Capitalize(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] >= 'A' && r[i] <= 'Z' {
			r[i] = r[i] + 32
		}
	}
	if r[0] >= 'a' && r[0] <= 'z' {
		r[0] = r[0] - 32
	}

	for i := 1; i < len(r); i++ {
		if !isAlphaNumeric(r[i-1]) {
			if r[i] >= 'a' && r[i] <= 'z' {
				r[i] = r[i] - 32
			} else {
				r[i] = r[i]
			}
		}
	}
	return string(r)
}

func isAlphaNumeric(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}
