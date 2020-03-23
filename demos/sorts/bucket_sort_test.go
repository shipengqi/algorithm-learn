package sorts

import (
	"testing"
)

func TestBucketSort(t *testing.T) {
	a := []int{1, 6, 3, 5, 8, 6, 4, 8, 10, 56, 43, 97, 26, 65, 43, 59, 75, 73, 91, 32, 17, 19, 72, 36, 31,}
	BucketSort(a)
	t.Log(a)
}
