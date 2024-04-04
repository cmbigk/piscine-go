package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a
	d := *b
	*a = c / d
	*b = c % d
}

// This function will divide the int a and b.
// The result of this division will be stored in the int pointed by a.
// The remainder of this division will be stored in the int pointed by b.
