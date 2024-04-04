package piscine

func Map(f func(int) bool, a []int) []bool {
	length := 0
	for index := range a {
		length = index + 1
	}
	Slice := make([]bool, length)

	for index := range a {
		Slice[index] = f(a[index])
	}

	return Slice
}
