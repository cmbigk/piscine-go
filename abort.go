package piscine

// Write a function named Abort that returns the median of five int arguments.

func Abort(a, b, c, d, e int) int {
	// arrange the numbers in ascending order
	numbers := []int{a, b, c, d, e}
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	// return the median
	if len(numbers)%2 == 0 {
		return (numbers[len(numbers)/2] + numbers[len(numbers)/2-1]) / 2
	} else {
		return numbers[len(numbers)/2]
	}
}
