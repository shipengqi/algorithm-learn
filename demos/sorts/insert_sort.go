package sorts

func InsertSort(a []int)  {
    length := len(a)
    if length <= 1 {
        return
    }
    for i := 1; i < length; i ++ {
        j := i - 1
        tmp := a[i]
        for ; j >= 0; j -- {
            if a[j] > tmp {
                a[j + 1] = a[j] // 移动数据
            } else {
                break
            }
        }
        a[j + 1] = tmp // 插入数据
    }
}