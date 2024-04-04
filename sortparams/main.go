package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var counter int
	for i := range os.Args {
		counter = i
	}
	for i := 1; i <= counter; i++ {
		for j := 1; j <= counter; j++ {
			if os.Args[i] < os.Args[j] {
				os.Args[i], os.Args[j] = os.Args[j], os.Args[i]
			}
		}
	}
	for i := range os.Args {
		if i > 0 {
			for _, item := range os.Args[i] {
				z01.PrintRune(item)
			}
			z01.PrintRune('\n')
		}
	}
}
