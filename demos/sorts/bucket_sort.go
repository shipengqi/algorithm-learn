package sorts

import "fmt"

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
	max := getMax(a)
	buckets := make([][]int, num)  // 二维切片，桶
	// 分桶
	index := 0
	for i := 0; i < num; i ++ {
		index = a[i] * (num - 1) / max  // 桶序号
		fmt.Printf("---------element: %d, index: %d-------------\n", a[i], index)
		buckets[index] = append(buckets[index], a[i]) // 加入对应的桶中
	}

	tmpPos := 0  // 标记数组位置
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(buckets[i])  // 桶内做快速排序
			copy(a[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
}