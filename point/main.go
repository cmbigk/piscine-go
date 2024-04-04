package main

import "github.com/01-edu/z01"

type point struct {
	x string
	y string
}

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

func main() {
	points := &point{}
	setPoint(points)
	str := "x = " + points.x + ", y = " + points.y
	printStr(str)
}

func printStr(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
