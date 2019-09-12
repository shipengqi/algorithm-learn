package sorts

func BubbleSort(a []int) {
	length := len(a)
	if len(a) <= 1 {
		return
	}
	for i := 0; i < length; i ++ {
		flag := false
		for j := 0; j < length - i - 1; j ++ {
			if a[j] > a[j + 1] {
				a[j], a[j + 1] = a[j + 1], a[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func InsertSort(a []int)  {
	length := len(a)
	if len(a) <= 1 {
		return
	}
	for i := 0; i < length; i ++ {
        value := a[i]
        j := i - 1
        for ; j >= 0; j -- {   // 移动
        	if a[j] > value {
        		a[j + 1] = a[j]
			} else {
				break
			}
		}
        a[j + 1] = value    // 插入
	}
}

func SelectionSort(a []int)  {
	length := len(a)
	if len(a) <= 1 {
		return
	}
	for i := 0; i < length; i ++ {
        minIndex := i
		for j := i + 1; j < length; j ++ { // 找到最小值
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i] // 交换
	}
}