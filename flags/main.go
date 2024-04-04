package main

import (
	"fmt"
	"os"
)

func main() {
	order := false
	ans := ""

	if len(os.Args) == 1 || len(os.Args) > 4 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Print("--insert\n")
		fmt.Print("  -i\n")
		fmt.Print("\t This flag inserts the string into the string passed as argument.\n")
		fmt.Print("--order\n")
		fmt.Print("  -o\n")
		fmt.Print("\t This flag will behave like a boolean, if it is called it will order the argument.\n")
		return
	}
	if len(os.Args) == 2 {
		fmt.Print(os.Args[1])
		return
	}
	ans = os.Args[len(os.Args)-1]

	if len(os.Args) == 3 {
		if os.Args[1] == "--order" || os.Args[1] == "-o" {
			order = true
		} else {
			insert(os.Args[1], &ans)
		}
	}

	if len(os.Args) == 4 {
		order = true
		if os.Args[1] == "--order" || os.Args[1] == "-o" {
			insert(os.Args[2], &ans)
		} else {
			insert(os.Args[1], &ans)
		}
	}

	if order {
		sort(&ans)
	}

	fmt.Print(ans, "\n")
}

func insert(arg string, s *string) {
	argArr := []rune(arg)
	for i := range argArr {
		if argArr[i] == '=' {
			for j := i + 1; j < len(argArr); j++ {
				*s += string(argArr[j])
			}
		}
	}
}

func sort(s *string) {
	sArr := []rune(*s)
	for i := range sArr {
		for j := i + 1; j < len(sArr); j++ {
			if sArr[i] > sArr[j] {
				sArr[i], sArr[j] = sArr[j], sArr[i]
			}
		}
	}
	*s = string(sArr)
}
