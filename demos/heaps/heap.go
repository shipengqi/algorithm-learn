package heaps

type Heap struct {
	a       []int  // 数组，从下标 1 开始存储
	n       int    // 堆可以存储的元素的最大个数
	count   int    // 堆已经存储的元素的个数
}

func NewHeap(capacity int) *Heap {
	return &Heap{make([]int, capacity + 1), capacity, 0}
}

// top-max heap -> heapify from down to up
func (h *Heap) insert(data int) {
    if h.count >= h.n { // 堆满了
		return
	}
    h.count ++
    h.a[h.count] = data

	//compare with parent node
	i := h.count
	parent := i / 2
	// fmt.Println(i, parent)
	for parent > 0 && h.a[parent] < h.a[i] {
		swap(h.a, parent, i)
		i = parent
		parent = i / 2
	}
}

func (h *Heap) removeMax()  {
	if h.count == 0 { // 堆空了
		return
	}

	// swap max and last
	swap(h.a, 1, h.count)

	// heapify from up to down
	heapifyUpToDown(h.a, h.count)
}

func heapifyUpToDown(a []int, count int) {
	for i := 1; i <= count/2; {
		maxIndex := i
		if a[i] < a[i * 2] {
			maxIndex = i * 2
		}

		if i * 2 + 1 <= count && a[maxIndex] < a[i * 2 + 1] {
			maxIndex = i * 2 + 1
		}

		if maxIndex == i {
			break
		}

		swap(a, i, maxIndex)
		i = maxIndex
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}