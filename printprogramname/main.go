package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	for _, r := range name {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
