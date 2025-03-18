---
title: 有序数组的平方
weight: 3
---

一个按**非递减顺序**排序的整数数组 nums，返回**每个数字的平方**组成的新数组，要求也按**非递减顺序**排序。

示例 ：

- 输入：`nums = [-4,-1,0,3,10]`
- 输出：`[0,1,9,16,100]`

## 实现

暴力的解法就是每个数平方之后，排个序

```go
import "sort"

func sortedSquares(nums []int) []int {
    for i, val := range nums {
        nums[i] *= val
    }
    sort.Ints(nums)
    return nums
}
```

## 双指针法（快慢指针法）

数组其实是有序的，只不过负数平方之后可能成为最大数了。

那么数组平方的最大值就在数组的两端，不是最左边就是最右边，不可能是中间。

双指针法，i 指向起始位置，j 指向终止位置。定义一个新数组 result，和 A 数组一样的大小，让 k 指向 result 数组终止位置。

```go
func sortedSquares(nums []int) []int {
	n := len(nums)
	var result = make([]int, n)
	i, j, k := 0, n-1, n-1
	for i <= j { // 注意这里要 i <= j，因为最后要处理两个元素
		if nums[i]*nums[i] > nums[j]*nums[j] {
			result[k] = nums[i] * nums[i]
			i++
		} else {
			result[k] = nums[j] * nums[j]
			j--
		}
		k--
	}
	return result
}
```