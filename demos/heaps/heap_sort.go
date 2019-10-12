package heaps

// build a heap
func buidHeap(a []int, n int) {

	// heapify from the last parent node
	for i := n / 2; i >= 1; i-- { // 对下标从 n/2 开始到 1 的数据进行堆化，因为叶子节点堆化比较的是自己，所以叶子节点不需要堆化
		heapifyUpToDown2(a, i, n) // 从第一个非叶子节点开始堆化
	}

}

// sort by ascend, a index begin from 1, has n elements
func sort(a []int, n int) {
	buidHeap(a, n)

	k := n
	for k >= 1 {
		swap(a, 1, k)
		heapifyUpToDown2(a, 1, k-1)
		k--
	}
}

// heapify from up to down , node index = top
func heapifyUpToDown2(a []int, top int, count int) {

	for i := top; i <= count/2; {

		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		swap(a, i, maxIndex)
		i = maxIndex
	}

}