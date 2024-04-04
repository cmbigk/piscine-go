package piscine

func ToUpper(s string) string {
	r := []rune(s)
	for i, c := range r {
		if c >= 'a' && c <= 'z' {
			r[i] = c - 32
		}
	}
	return string(r)
}
