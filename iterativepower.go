package piscine

func IterativePower(nb, p int) int {
	if p < 0 {
		return 0
	} else if p == 1 {
		return nb
	} else if p == 0 {
		return 1
	} else {
		result := 1
		for i := 0; i < p; i++ {
			result *= nb
		}
		return result
	}
}
