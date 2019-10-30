package sma

import "fmt"

// bc: pattern char index hash mapping
func GenerateBC(pattern string) []int {

	bc := make([]int, 256)

	for index := range bc {
		bc[index] = -1
	}

	for index, char := range pattern {
		fmt.Println(char)
		fmt.Println(int(char))
		bc[int(char)] = index
	}

	return bc
}