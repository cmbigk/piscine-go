package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	table := os.Args[1:]
	runes := make([]byte, len(table))
	for k := 0; k < len(table); k++ {
		runes[k] = byte((table[k])[0])
		if len(table[k]) == 2 && runes[k] == 49 {
			runes[k] = 10 + table[k][1]
		}
		if len(table[k]) == 2 && runes[k] == 50 {
			runes[k] = 20 + table[k][1]
		}
		if len(table[k]) == 2 && runes[k] == 51 {
			runes[k] = 30 + table[k][1]
		}
		if len(table[k]) == 2 && runes[k] == 52 {
			runes[k] = 40 + table[k][1]
		}
		if len(table[k]) == 2 && runes[k] == 53 {
			runes[k] = 50 + table[k][1]
		}
		if len(table[k]) == 2 && runes[k] == 54 {
			runes[k] = 60 + table[k][1]
		}
		if len(table[k]) == 2 && runes[k] == 55 {
			runes[k] = 70 + table[k][1]
		}
		if len(table[k]) == 2 && runes[k] == 56 {
			runes[k] = 80 + table[k][1]
		}
		if len(table[k]) == 2 && runes[k] == 57 {
			runes[k] = 90 + table[k][1]
		}
	}
	if len(runes) == 0 {
	} else if runes[0] == 45 {
		for j := range runes {
			if runes[j] > 48 && runes[j] <= 48+26 {
				runes[j] += 16
			} else {
				runes[j] = ' '
			}
		}
		for i := 1; i < len(runes); i++ {
			z01.PrintRune(rune(runes[i]))
		}
		z01.PrintRune('\n')
	} else {
		for j := range runes {
			if runes[j] > 48 && runes[j] <= 48+26 {
				runes[j] += 48
			} else {
				runes[j] = ' '
			}
		}
		for i := 0; i < len(runes); i++ {
			z01.PrintRune(rune(runes[i]))
		}
		z01.PrintRune('\n')
	}
}
