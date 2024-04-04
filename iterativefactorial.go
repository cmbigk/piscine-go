package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 25 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= nb; i++ {
		result *= i // result = result * i
	}
	return result
}

// first make a for loop starting from highest valoue to lowest value

// I need to multiply the current value of i/index (which is my number), with the number smaller than that
// i * (i - 1)
// I should avoid multiplying by 0, so I will end ther multiplication when the number (i -1) reaches 1.
