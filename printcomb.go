package main

import "github.com/01-edu/z01"

func main() {
	for a := '0'; a <= '7'; a++ {
		
		for b := '1'; b <= '8'; b++ {
			
			for c := '2'; c <= '9'; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				z01.PrintRune(10)
			} 
		}
	}
	
	z01.PrintRune(44)
	z01.PrintRune(32)
}
