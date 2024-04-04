package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for c1 := '0'; c1 <= '9'; c1++ {
		for c2 := '0'; c2 <= '9'; c2++ {
			for c3 := '0'; c3 <= '9'; c3++ {
				for c4 := '0'; c4 <= '9'; c4++ {
					if c1 < c3 || c1 == c3 && c2 < c4 {
						z01.PrintRune(c1)
						z01.PrintRune(c2)
						z01.PrintRune(' ')
						z01.PrintRune(c3)
						z01.PrintRune(c4)
						if c1 == '9' && c2 == '8' && c3 == '9' && c4 == '9' {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}

					}
				}
			}
		}
	}
}
