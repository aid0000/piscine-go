package piscine

import "fmt"

func PrintStr(a string) {
	for _, v := range a {
		fmt.Print(string(v))
	}
}
