package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var nb []int
	for n > 0 {
		nb = append(nb, n%10)
		n = n / 10
	}
	for i := 0; i < len(nb)-1; i++ {
		for j := i + 1; j < len(nb); j++ {
			if nb[i] > nb[j] {
				nb[i], nb[j] = nb[j], nb[i]
			}
		}
	}
	for _, nb := range nb {
		z01.PrintRune(rune(nb + '0'))
	}
}
