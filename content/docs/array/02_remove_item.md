---
title: 移除元素
weight: 2
---

一个数组 nums 和一个值 val，需要**原地**移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，必须仅使用 O(1) 额外空间并原地修改输入数组。

## 实现

暴力的解法就是两层 for 循环，一个 for 循环遍历数组元素 ，第二个 for 循环更新数组。

```go
func removeElement(nums []int, val int) int {
    size := len(nums)
    for i := 0; i < size; i++ {
		if nums[i] == val { // 发现需要移除的元素，就将数组集体向前移动一位
			for j := i + 1; j < size; j++ {
				nums[j-1] = nums[j]
			}
			i-- // 因为下标 i 以后的数值都向前移动了一位，所以 i 也向前移动一位	
			size-- // 此时数组的大小-1
		}
	}
	return size
}
```

## 双指针法（快慢指针法）

通过一个快指针和慢指针在一个 for 循环下完成两个 for 循环的工作。

```go
func removeElement(nums []int, val int) int {
    slow := 0
    for fast := 0; fast < size; fast++ {
		if nums[fast] != val { 
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
```