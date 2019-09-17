package sorts

// 获取数组中最大值
func getMax(a []int) int {
	max := a[0]
	for i := 0; i < len(a); i ++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func BucketSort(a []int)  {
	num := len(a)
	if num <= 1{
		return
	}
}