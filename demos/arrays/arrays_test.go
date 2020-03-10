package arrays

import "testing"

func TestArray_Insert(t *testing.T) {
    t.Run("error: full array", func(t *testing.T) {
        arr := NewArray(3)
        want := "full array"
        _ = arr.Insert(0, 1)
        _ = arr.Insert(1, 2)
        _ = arr.Insert(2, 3)
        err := arr.Insert(3, 4)
        if want != err.Error() {
            t.Fatalf("want %s, got %s", want, err.Error())
        }
        arr.Print()
    })

    t.Run("error: index out of range", func(t *testing.T) {
        arr := NewArray(3)
        want := "index out of range"
        err := arr.Insert(3, 1)
        if want != err.Error() {
            t.Fatalf("want %s, got %s", want, err.Error())
        }
        arr.Print()
    })
}

func TestArray_Find(t *testing.T) {
    t.Run("error: full array", func(t *testing.T) {
        arr := NewArray(3)
        want := 2
        _ = arr.Insert(0, 1)
        _ = arr.Insert(1, 2)
        _ = arr.Insert(2, 3)
        v, _ := arr.Find(1)
        if want != v {
            t.Fatalf("want %d, got %d", want, v)
        }
        arr.Print()
    })

    t.Run("error: index out of range", func(t *testing.T) {
        arr := NewArray(3)
        want := "index out of range"
        _, err := arr.Find(3)
        if want != err.Error() {
            t.Fatalf("want %s, got %s", want, err.Error())
        }
        arr.Print()
    })
}

func TestArray_Delete(t *testing.T) {
    t.Run("error: index out of range", func(t *testing.T) {
        arr := NewArray(3)
        want := "index out of range"
        err := arr.Insert(3, 1)
        if want != err.Error() {
            t.Fatalf("want %s, got %s", want, err.Error())
        }
        arr.Print()
    })

    t.Run("length 2", func(t *testing.T) {
        arr := NewArray(3)
        want := 2
        _ = arr.Insert(0, 1)
        _ = arr.Insert(1, 2)
        _ = arr.Insert(2, 3)
        v, _ := arr.Delete(1)
        if want != v {
            t.Fatalf("want %d, got %d", want, v)
        }
        arr.Print()
    })
}

func TestArray_Append(t *testing.T) {
    t.Run("length 2", func(t *testing.T) {
        arr := NewArray(3)
        want := 2
        _ = arr.Insert(0, 1)
        _ = arr.Append(2)
        if uint(want) != arr.Len() {
            t.Fatalf("want %d, got %d", want, arr.Len())
        }
        arr.Print()
    })
}