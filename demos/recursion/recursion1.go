package recursion

import "fmt"

// 递归实现斐波那契数列 F(1)=1，F(2)=1, F(n)=F(n-1)+F(n-2)（n>=3，n∈N*）
type Fibs struct {
	val map[int]int // 使用字典存储结果
}

func NewFibs(n int) *Fibs {
	return &Fibs{make(map[int]int, n)}
}

func (f *Fibs) Fibonacci(n int) int {
	if f.val[n] != 0 {
		return f.val[n]
	}
	if n <= 1 {
		f.val[1] = 1
		return 1
	} else if n == 2 {
		f.val[2] = 1
		return 2
	} else {
		result := f.Fibonacci(n - 1) + f.Fibonacci(n - 2)
		f.val[n] = result
		return result
	}
}

func (f *Fibs) Print(n int) {
	fmt.Println(f.val[n])
}