package sorts

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{3, 5, 4, 2, 6, 7, 9, 1, 8}
	BubbleSort(arr)
	t.Log(arr)
	t.Run("BubbleSort", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(arr, want) {
			t.Fatalf("Wanted %v, got %v", want, arr)
		}
	})
}

func TestInsertSort(t *testing.T) {
	arr := []int{3, 5, 4, 2, 6, 7, 9, 1, 8}
	BubbleSort(arr)
	t.Log(arr)
	t.Run("InsertSort", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(arr, want) {
			t.Fatalf("Wanted %v, got %v", want, arr)
		}
	})
}

func TestSelectionSort(t *testing.T) {
	arr := []int{3, 5, 4, 2, 6, 7, 9, 1, 8}
	SelectionSort(arr)
	t.Log(arr)
	t.Run("InsertSort", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(arr, want) {
			t.Fatalf("Wanted %v, got %v", want, arr)
		}
	})
}
