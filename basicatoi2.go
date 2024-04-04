package piscine

func BasicAtoi2(s string) int {
	sign := 1
	var num int
	startIndex := 0

	if s == "" || s == "0" {
		return 0
	}

	if s[0] == '-' {
		sign = -1
		startIndex = 1
	}
	if s[0] == '+' {
		startIndex = 1
	}

	for i := startIndex; i <= len(s)-1; i++ {
		if rune(s[i]) < '0' || rune(s[i]) > '9' {
			return 0
		}
		// fmt.Printf("%T => %v |  ", int(rune(s[i])-'0'), int(rune(s[i])-'0'))
		num = num*10 + int(rune(s[i])-'0')
	}

	return num * sign
}

// TESTING
// func main() {
// 	fmt.Println(BasicAtoi(""))
// 	fmt.Println(Atoi("12345"))
// 	fmt.Println(Atoi("0000000012345"))
// 	fmt.Println(Atoi("012 345"))
// 	fmt.Println(Atoi("Hello World!"))
// 	fmt.Println(Atoi("+1234"))
// 	fmt.Println(Atoi("-1234"))
// 	fmt.Println(Atoi("++1234"))
// 	fmt.Println(Atoi("--1234"))
// }

// func main() {
// 	fmt.Println(piscine.Atoi("12345"))            =>   12345
// 	fmt.Println(piscine.Atoi("0000000012345"))    =>   12345
// 	fmt.Println(piscine.Atoi("012 345"))          =>   0
// 	fmt.Println(piscine.Atoi("Hello World!"))     =>   0
// 	fmt.Println(piscine.Atoi("+1234"))            =>   1234
// 	fmt.Println(piscine.Atoi("-1234"))            =>   -1234
// 	fmt.Println(piscine.Atoi("++1234"))           =>   0
// 	fmt.Println(piscine.Atoi("--1234"))           =>   0
// }

// OUTPUT :
// 12345
// 12345
// 0
// 0
// 1234
// -1234
// 0
// 0
// $
