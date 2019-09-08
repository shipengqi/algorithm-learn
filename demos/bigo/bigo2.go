package main

import "fmt"

func cal3(n int) int {
	ret := 0
	for i := 1; i <= n; i++ {
		ret = ret + f(i)
	}
	return ret
}

func f(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum
}

func DemopPint(n int) {
	a := make([]int, n)
	for i := 0; i < n; i ++ {
		a = append(a, i * i)
	}
	for i := n - 1; i >= 0; i -- {
		fmt.Println(a[i])
	}
}

func find(array []int, n, x int) int {
	pos := -1
	for i := 0; i < n; i ++ {
		if array[i] == x {
			pos = i
		}
	}
	return pos
}

func main() {
	DemopPint(5)
}