---
title: 螺旋矩阵
weight: 5
---

给定一个正整数 n ，生成一个包含 1 到 `n^2` 所有元素，且元素按顺时针顺序螺旋排列的 `n x n` 正方形矩阵 matrix 。

示例 ：

- 输入：3
- 输出：`[ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]`


## 实现

在解决螺旋矩阵问题时，使用循环不变量规则可以更好地控制边界条件和循环逻辑。**循环不变量规则指的是在循环过程中保持不变的属性或条件**。

模拟顺时针画矩阵的过程:

- 填充上行从左到右
- 填充右列从上到下
- 填充下行从右到左
- 填充左列从下到上

由外向内一圈一圈这么画下去。

可以发现这里的边界条件非常多，在一个循环中，如此多的边界条件，一定要坚持循环不变量规则。

```go
func generateMatrix(n int) [][]int {
	startx, starty := 0, 0 // 起始位置
	var loop int = n / 2 // 旋转圈数
	var center int = n / 2 // 中间位置
	count := 1 // 更新填充数字
	offset := 1 // 控制每一层填充元素个数
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for loop > 0 {
		i, j := startx, starty

		// 行数不变 列数在变
		for j = starty; j < n-offset; j++ {
			res[startx][j] = count
			count++
		}
		// 列数不变是 j 行数变
		for i = startx; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		// 行数不变 i 列数变 j--
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		// 列不变 行变
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}
		startx++
		starty++
		offset++
		loop--
	}
    // 如果 n 为奇数的话，需要单独给矩阵最中间的位置赋值
	if n%2 == 1 {
		res[center][center] = n * n
	}
	return res
}
```