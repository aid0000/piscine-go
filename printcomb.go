package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := 0; a <= 7; a++ {
		z01.PrintRune(a)
		for b := (a + 1); b <= 8; b++ {
			z01.PrintRune(b)
			for c := b + 1; c <= 9; c++ {
				z01.PrintRune(c)
			}
		}
	}
	z01.PrintRune(44)
	z01.PrintRune(32)
}
