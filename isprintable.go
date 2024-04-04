package piscine

func IsPrintable(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c < ' ' || c > '~' {
			return false
		}
	}
	return true
}
