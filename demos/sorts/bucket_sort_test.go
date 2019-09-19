package sorts

import (
	"testing"
)

func TestBucketSort(t *testing.T) {
	a := []int{1, 6, 3, 5, 8, 6, 4, 8, 10, 56, 43, 97, 26, 65, 43, 59, 75, 73, 91, 32, 17, 19, 72, 36, 31,}
	BucketSort(a)
	t.Log(a)
	//t.Run("BucketSort", func(t *testing.T) {
	//	want := []int{1, 3, 4, 5, 6, 6, 8}
	//	if !reflect.DeepEqual(a, want) {
	//		t.Fatalf("Wanted %v, got %v", want, a)
	//	}
	//})
}
