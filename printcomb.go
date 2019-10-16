package piscine

import ( "github.com/01-edu/z01" )

func PrintComb2() {
	for a := 0; a < 10; a=a+1 {
		for b := 0 ; b < 10; b= b+1 {
			for c := a; c < 10; c= c+1{
				for d := d < 10; d = d + 1; d= b + 1; d := 0 {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)
					if a < 9 || b < 8 || c < 9 || d < 9{
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}

		}
	}
	z01.PrintRune(10)
}
