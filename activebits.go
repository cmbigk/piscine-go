package piscine

func ActiveBits(n int) (s int) {
	for ; n > 1; n /= 2 {
		s += n % 2
	}
	s += n
	return
}
