package sorts

func CountSort( a []int)  {
	num := len(a)
	if num <= 1{
		return
	}
	// 查找数组中数据的范围
	var max int = a[0]
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}

	// c 为计数数组 下标 [0, max]
	c := make([]int, max + 1)

	// 计算每个元素的个数，放入 c 中
	for i := range a {
		c[a[i]] ++
	}

	// 依次累加
	for i := 1; i <= max; i ++ {
		c[i] += c[i-1]
	}

	// 临时数组 r，存储排序之后的结果
	r := make([]int, num)
	for i := range a {
		index := c[a[i]] - 1
		r[index] = a[i]
		c[a[i]]--
	}

	copy(a, r)
}
