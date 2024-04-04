package piscine

func ConvertBase(num, baseFrom, baseTo string) string {
	return FromDigital2CustomBase(AtoiBase(num, baseFrom), baseTo)
}

func FromDigital2CustomBase(nbr int, baseTo string) string {
	result := ""
	if !validBase(baseTo) {
		return result
	}
	if baseTo == "0123456789" {
		return iota(nbr)
	}
	baseRunes := []rune(baseTo)
	var digits []int
	for nbr > 0 {
		digits = append(digits, nbr%len(baseRunes))
		nbr = nbr / len(baseRunes)
	}
	for i := len(digits) - 1; i >= 0; i-- {
		result += string(baseRunes[digits[i]])
	}
	return result
}

func validBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	baseRunes := []rune(base)
	for i := 0; i < len(baseRunes)-1; i++ {
		if baseRunes[i] == '+' || baseRunes[i] == '-' || baseRunes[i+1] == '+' || baseRunes[i+1] == '-' { // A base should not contain + or - characters.
			return false
		}
		for j := i + 1; j < len(baseRunes); j++ {
			if baseRunes[i] == baseRunes[j] {
				return false
			}
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	var i, j int
	var result int
	var valid bool
	var baseLen int
	var sLen int
	var sArray []rune
	var baseArray []rune

	baseLen = 0
	sLen = 0
	result = 0
	valid = true

	for i = range base {
		baseLen++
	}

	for i = range s {
		sLen++
	}

	if baseLen < 2 {
		valid = false
	}

	for i = range base {
		for j = range base {
			if i != j && base[i] == base[j] {
				valid = false
			}
		}
	}

	if base[0] == '+' || base[0] == '-' {
		valid = false
	}

	if !valid {
		return 0
	}

	sArray = []rune(s)
	baseArray = []rune(base)

	for i = range sArray {
		for j = range baseArray {
			if sArray[i] == baseArray[j] {
				result = result*baseLen + j
			}
		}
	}

	return result
}
