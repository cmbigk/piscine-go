package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 || s == "" {
		return true
	}

	for _, v := range s {
		if (v < 'a' || v > 'z') && (v < '0' || v > '9') && (v < 'A' || v > 'Z') {
			return false
		}
	}
	return true
}
