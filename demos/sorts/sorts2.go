package sorts

func MergeSort(a []int)  {
	length := len(a)
	if len(a) <= 1 {
		return
	}
	mergeSort(a, 0, length - 1)
}

func mergeSort(arr []int, start, end int)  {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid + 1, end)
	merge(arr, start, mid, end)
}

// arr 是原切片，只传入分组后的下标
func merge(arr []int, start, mid, end int) {
	// 申请一个可以存放两个分组所有元素的临时切片
	tmpArr := make([]int, end - start + 1)
	i := start // 第一个分组的开始位置
	j := mid + 1 // 第二个分组的开始位置
	k := 0
	for ; i <= mid && j <= end; k++ { // 每个分组都已经排好顺序
		if arr[i] < arr[j] {  // 直接比较每个分组相同位置的元素大小
			tmpArr[k] = arr[i] // 小的元素放到临时切片中
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}
	copy(arr[start : end + 1], tmpArr) // 将临时切片的元素拷贝到原切片的对应位置
}

func QuickSort(a []int) {
	length := len(a)
	if len(a) <= 1 {
		return
	}
	separateSort(a, 0, length - 1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	// 选取最后一位当对比数字
	pivot := arr[end]

	var i = start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if !(i == j) {
				// 交换位置
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	return i
}