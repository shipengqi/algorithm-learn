package main

import "fmt"

func cal(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum
}

func cal2(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j ++ {
			sum = sum + i
		}
	}
	return sum
}
func main() {
	r := cal(10)
	r2 := cal2(10)
	fmt.Println(r, r2)
}
