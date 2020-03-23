package sorts

func SelectionSort(a []int)  {
    length := len(a)
    if len(a) <= 1 {
        return
    }
    for i := 0; i < length; i ++ {
        minIndex := i
        j := i + 1
        for ; j < length; j ++ {
            if a[j] < a[minIndex] {
                minIndex = j // 找到最小值下标
            }
        }
        // 交换数据
        a[i], a[minIndex] = a[minIndex], a[i]
    }
}
