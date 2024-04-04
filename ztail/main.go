package main

import (
	"fmt"
	"os"
)

var errors = false

func main() {
	if len(os.Args) == 4 {
		readFile(os.Args[3], true, true)
	} else if len(os.Args) > 4 {
		for i, file := range os.Args[3:] {
			if i == 0 {
				readFile(file, false, true)
			} else {
				readFile(file, false, false)
			}
		}
	}
	if errors {
		os.Exit(1)
	}
}

func BasicAtoi(s string) int {
	sArr := []rune(s)
	n := 0
	for i := 0; i < len(sArr); i++ {
		n *= 10
		n += int(sArr[i] - '0')
	}
	return n
}

func readFile(filename string, solo, first bool) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err.Error() + "\n")
		errors = true
		return
	}
	if !first {
		fmt.Print("\n")
	}
	if !solo {
		fmt.Print("==> " + string(filename) + " <==\n")
	}
	start := len(data) - BasicAtoi(os.Args[2])
	if start < 0 {
		start = 0
	}
	fmt.Print(string(data[start:]))
}
