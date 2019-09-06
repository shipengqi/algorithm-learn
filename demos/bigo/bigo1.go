package main

import "fmt"

func cal(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum
}

func main() {
	r := cal(10)
	fmt.Println(r)
}
