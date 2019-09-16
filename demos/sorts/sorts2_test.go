package sorts

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{3, 5, 4, 2, 6, 7, 9, 1, 8}
	BubbleSort(arr)
	t.Log(arr)
	t.Run("MergeSort", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(arr, want) {
			t.Fatalf("Wanted %v, got %v", want, arr)
		}
	})
}

func TestQuickSort(t *testing.T) {
	arr := []int{3, 5, 4, 2, 6, 7, 9, 1, 8}
	BubbleSort(arr)
	t.Log(arr)
	t.Run("QuickSort", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(arr, want) {
			t.Fatalf("Wanted %v, got %v", want, arr)
		}
		if want[0] != 1 {
			t.Fatalf("Wanted %v, got %v", 1, want[0])
		}
	})
}
