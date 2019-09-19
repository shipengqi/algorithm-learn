package sorts

func BinarySearch(a []int, value int) int {
	length := len(a)
	if length < 1 {
		return -1
	}
	start := 0
	end := length -1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == value {
			return mid
		} else if a[mid] < value {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func BinarySearchRecursive(a []int, value int) int {
	length := len(a)
	if length < 1 {
		return -1
	}

	return bs(a, value, 0, length -1)
}

func bs(a []int, value, start, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if a[mid] == value {
		return mid
	} else if a[mid] < value {
		return bs(a, value, mid + 1, end)
	} else {
		return bs(a, value, start, mid - 1)
	}
}

