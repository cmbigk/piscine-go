package piscine

// ConcatParams([]string{"Hello", "how", "are", "you?"})
// your output: "Hello\nhow\nare\nyou?\n"
// expected: "Hello\nhow\nare\nyou?"

func ConcatParams(args []string) string {
	ans := ""
	for i, v := range args {
		if i == len(args)-1 {
			ans += v
			break
		}
		ans += (v + string('\n'))
	}
	return ans
}
