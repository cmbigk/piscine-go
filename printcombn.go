package piscine

import "github.com/01-edu/z01"

func intToString(num int) string {
	return string(num + '0')
}

func iota(num int) string {
	var result string
	if num == 0 {
		return "0"
	}
	for num > 0 {
		result = intToString(num%10) + result
		num /= 10
	}
	return result
}

func printString(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
}

func printComb(n int, prev int, result string, count *int) {
	for i := 0; i < 10; i++ {
		if prev < i {
			if n == 1 {
				if *count > 0 {
					printString(", ")
				}
				printString(result + iota(i))
				*count++
			} else {
				printComb(n-1, i, result+iota(i), count)
			}
		}
	}
}

func PrintCombN(n int) {
	var count int
	for i := 0; i < 10; i++ {
		if n > 1 {
			printComb(n-1, i, iota(i), &count)
		} else {
			if count > 0 {
				printString(", ")
			}
			printString(iota(i))
			count++
		}
	}
	printString("\n")
}
