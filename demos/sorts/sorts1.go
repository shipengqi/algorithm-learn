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

}