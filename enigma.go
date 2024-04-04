package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	x := *******c   // store the value of c before it gets lost
	*******c = ***a // a into c. DONE
	y := ****d
	****d = x // c into d. DONE
	z := *b
	*b = y
	***a = z
}
