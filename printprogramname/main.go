package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	PrintProgramName()
	z01.PrintRune('\n')
}

func PrintProgramName() {
	baseName := ""
	if len(os.Args) > 0 {
		baseName = os.Args[0]
	}
	nameOfProgram := ""

	if len(baseName) > 0 {
		for i := len(baseName) - 1; i >= 0; i-- {
			if baseName[i] == '/' {

				nameOfProgram = baseName[i+1:]
				break
			}
		}
		if nameOfProgram == "" {
			nameOfProgram = baseName
		}
	}
	for _, letter := range nameOfProgram {
		z01.PrintRune(letter)
	}
}
