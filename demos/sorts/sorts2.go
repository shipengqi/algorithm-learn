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

func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end - start + 1)
	copy(arr[start : end + 1], tmpArr)
}
