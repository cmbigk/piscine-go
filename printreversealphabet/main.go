package main

import "github.com/01-edu/z01"

func main() {
	for y := 'z'; y >= 'a'; y-- {
		z01.PrintRune(y)
	}
	z01.PrintRune('\n')
}
