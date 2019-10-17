package piscine

import "github.com/01-edu/z01"

func PrintStr(a string) {
	for _, v := range a {
		fmt.Print(string(v))
	}
}
