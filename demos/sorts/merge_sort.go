package sorts

func MergeSort(a []int)  {
    length := len(a)
    if length <= 1 {
        return
    }
    mergeSort(a, 0, length - 1)
}

func mergeSort(a []int, start, end int)  {
    if start >= end {
        return
    }
    p := (end + start) / 2
    mergeSort(a, p + 1, end)
    mergeSort(a, start, p)
    merge(a, start, p, end)
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