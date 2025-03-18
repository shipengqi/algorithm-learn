---
title: 二分查找
weight: 1
---

二分查找是一种非常简单易懂的快速查找算法。

## 实现

二分查找逻辑比较简单，但是涉及的很多的边界条件。例如到底是 `while(left < right)` 还是 `while(left <= right)`，到底是 `right = middle` ，还是要 `right = middle - 1` ？

```go
func search(nums []int, target int) int {
	length := len(nums)
	if length < 1 {
		return -1
	}
    var mid = 0
	// 定义 target 在左闭右闭的区间里，[left, right]
    var left = 0
    var right = length - 1
    for left <= right { // 当 left == right，区间 [left, right] 依然有效，所以用 <=
        if (nums[left] == target) {
            return left
        }
        if (nums[right] == target) {
            return right
        }
        mid = (left + right) / 2
        if (target < nums[mid]) {
            right = mid - 1 // 赋值为 mid - 1，因为当前这个 nums[mid] 一定不是 target，那么接下来要查找的左区间结束下标位置就是 mid - 1
        } else if (target > nums[mid]) {
            left = mid + 1  // 和 mid - 1 一样的道理
        } else {
			return mid
		}
    }
    return -1
}
```

递归实现二分查找：

```go
func search(nums []int, target int) int {
	length := len(nums)
	if length < 1 {
		return -1
	}

	return bs(nums, target, 0, length - 1)
}

func bs(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}
	if (nums[left] == target) {
        return left
    }
    if (nums[right] == target) {
        return right
    }
	mid := (left + right) / 2
	if nums[mid] < target {
		return bs(nums, target, mid + 1, right)
	} else if nums[mid] > target {
		return bs(nums, target, left, mid - 1)
	} else {
		return mid
	}
}
```

{{< callout type="info" >}}
上面的代码中，`mid := (left + right) / 2` 这种写法是有问题的。因为如果 left 和 right 比较大的话，两者之和就有可能会溢出。
- 改进的方法是将 mid 的计算方式写成 `left + (left - right) / 2`。
- 进一步优化性能，可以将除以 2 转化成位运算 `left + ((right - left) >> 1)`。相比除法运算来说，计算机处理位运算要快得多。
{{< /callout >}}


## 查找第一个值等于给定值的元素

前面实现的二分查找，适用于有序数据集合中不存在重复的数据的情况，如果有重复数据，如何找到第一个等于给定值的元素？

```go
func search(nums []int, target int) int {
	length := len(nums)
	if length < 1 {
		return -1
	}
	left := 0
	right := length -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if mid == 0 || nums[mid - 1] != target {
				return mid
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
```

注意 `nums[mid] == target` 时，要判断否有重复值，数据集合是有序的，要找到第一个等于 target 得元素：
- 如果 `mid == 0` 说明，已经是 nums 集合的第一个元素，不需要再继续查找。
- 如果 `nums[mid - 1] != target`，也就是 `nums[mid]` 的前一个元素不等于 target `nums[mid]` 就是第一个等于 target 得元素。

## 查找最后一个值等于给定值的元素

```go
func search(nums []int, target int) int {
	length := len(nums)
	if length < 1 {
		return -1
	}
	left := 0
	right := length -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if mid == length - 1 || nums[mid + 1] != target {
				return mid
			} else {
				right = mid + 1
			}
		}
	}
	return -1
}
```

和查找最后一个值等于给定值的元素思路差不多，注意 `nums[mid] == target` 时：
- 如果 `mid == length - 1` 说明，已经是 nums 集合的最后一个元素，不需要再继续查找。
- 如果 `nums[mid + 1] != target`，也就是 `nums[mid]` 的下一个元素不等于 target，那么就说明 `nums[mid]` 就是最后一个等于 target 得元素。