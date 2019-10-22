package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	intToMass(n)
}

func intToMass(n int) {
	var ResultMass []int
	for n != 0 {
		ResultMass = append(ResultMass, n%10)
		n = n / 10
	}
	sortMass(ResultMass)
}

func sortMass(ResultMass []int) {
	var temp int
	index := 0
	for _, key := range ResultMass {
		index++
		key = key
	}
	for i := 0; i < index-1; i++ {
		for j := 0; j < index-i-1; j++ {
			if ResultMass[j] < ResultMass[j+1] {
				temp = ResultMass[j]
				ResultMass[j] = ResultMass[j+1]
				ResultMass[j+1] = temp
			}
		}
	}
	printMass(ResultMass)
}

func printMass(ResultMass []int) {
	for _, value := range ResultMass {
		z01.PrintRune(rune(value) + '0')
	}
}
