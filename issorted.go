package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := false
	descending := false
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			ascending = true
		} else if f(a[i], a[i+1]) < 0 {
			descending = true
		}
	}
	if ascending && !descending {
		return true
	}
	if descending && !ascending {
		return true
	}
	if ascending && descending {
		return false
	}

	return true
}
