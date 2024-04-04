package piscine

func Max(a []int) int {
	var max int
	for i := 0; i < len(a); i++ {
		if i == 0 {
			max = a[i]
		} else if a[i] > max {
			max = a[i]
		}
	}
	return max
}
