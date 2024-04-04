package piscine

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}
