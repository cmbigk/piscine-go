package piscine

func FindNextPrime(nb int) int {
	for i := nb; true; i++ {
		if IsPrime1(i) {
			return i
		}
	}
	return 0
}

func IsPrime1(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
