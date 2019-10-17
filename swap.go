package piscine

import "fmt"

func Swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
	fmt.Println(a)
	fmt.Println(b)
}
