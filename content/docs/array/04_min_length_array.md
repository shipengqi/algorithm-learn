---
title: 长度最小的子数组
weight: 4
---

一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其**总和大于等于** target 的**长度最小**的**子数组** `[numsl, numsl+1, ..., numsr-1, numsr]` ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 ：

- 输入：`target = 7`, `nums = [2,3,1,2,4,3]`
- 输出：2
- 解释：子数组 `[4,3]` 是该条件下的长度最小的子数组。

## 实现

暴力的解法就是两个 for 循环，然后不断的寻找符合条件的子序列，时间复杂度很明显是 `O(n^2)`。

```go
func minSubArrayLen(target int, nums []int) int {
	result, sum, subLength := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target { // 一旦发现子序列和超过了 target，更新 result
				subLength = j - i + 1 // 取子序列的长度
				if result == 0 || result > subLength { // 更短的子序列，更新 result
					result = subLength
				}
				break
			}
		}
    }
	return result
}
```

## 滑动窗口

所谓滑动窗口，就是**不断的调节子序列的起始位置和终止位置，从而得出我们要想的结果**。滑动窗口也可以理解为双指针法的一种，只不过这种解法更像是一个窗口的移动，所以叫做滑动窗口更适合一些。

```go
func minSubArrayLen(target int, nums []int) int {
    i := 0
    l := len(nums)  // 数组长度
    sum := 0        // 子数组之和
    result := l + 1 // 初始化返回长度为 l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况
    for j := 0; j < l; j++ {
        sum += nums[j]
        for sum >= target {
            subLength := j - i + 1
            if subLength < result {
                result = subLength
            }
            sum -= nums[i]
            i++
        }
    }
    if result == l+1 {
        return 0
    } else {
        return result
    }
}
```