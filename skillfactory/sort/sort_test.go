package sort

import (
	"math/rand"
	"testing"
)

func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}

func TestSort_Bubble(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{[]int{5, 3, 4, 8, -12, 10}},
		{[]int{2, 4, 7, 1, 2}},
		{[]int{-10, -13, 12, 4, 3, 9}},
	}

	ascendingOrder := func(in []int) bool {
		prev := in[0]
		for _, val := range in[1:] {
			if val < prev {
				return false
			}
			prev = val
		}

		return true
	}

	for _, tt := range testCases {
		bubbleSort(tt.input)
		if !ascendingOrder(tt.input) {
			t.Error("order is not valid")
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.ReportAllocs()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
}

func TestSort_Selection(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{[]int{5, 3, 4, 8, -12, 10}},
		{[]int{2, 4, 7, 1, 2}},
		{[]int{-10, -13, 12, 4, 3, 9}},
	}

	ascendingOrder := func(in []int) bool {
		prev := in[0]
		for _, val := range in[1:] {
			if val < prev {
				return false
			}
			prev = val
		}

		return true
	}

	for _, tt := range testCases {
		selectionSort(tt.input)
		if !ascendingOrder(tt.input) {
			t.Error("order is not valid")
		}
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	b.ReportAllocs()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})
}

func TestSort_Insertion(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{[]int{5, 3, 4, 8, -12, 10}},
		{[]int{2, 4, 7, 1, 2}},
		{[]int{-10, -13, 12, 4, 3, 9}},
	}

	ascendingOrder := func(in []int) bool {
		prev := in[0]
		for _, val := range in[1:] {
			if val < prev {
				return false
			}
			prev = val
		}

		return true
	}

	for _, tt := range testCases {
		insertionSort(tt.input)
		if !ascendingOrder(tt.input) {
			t.Error("order is not valid")
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.ReportAllocs()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})
}

func TestSort_Merge(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{[]int{5, 3, 4, 8, -12, 10}},
		{[]int{2, 4, 7, 1, 2}},
		{[]int{-10, -13, 12, 4, 3, 9}},
	}

	ascendingOrder := func(in []int) bool {
		prev := in[0]
		for _, val := range in[1:] {
			if val < prev {
				return false
			}
			prev = val
		}

		return true
	}

	for _, tt := range testCases {
		mergeSort(tt.input)
		if !ascendingOrder(tt.input) {
			t.Error("order is not valid")
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	b.ReportAllocs()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			mergeSort(ar)
			b.StopTimer()
		}
	})
}

func TestSort_Quick(t *testing.T) {
	testCases := []struct {
		input []int
	}{
		{[]int{5, 3, 4, 8, -12, 10}},
		{[]int{2, 4, 7, 1, 2}},
		{[]int{-10, -13, 12, 4, 3, 9}},
	}

	ascendingOrder := func(in []int) bool {
		prev := in[0]
		for _, val := range in[1:] {
			if val < prev {
				return false
			}
			prev = val
		}

		return true
	}

	for _, tt := range testCases {
		quickSort(tt.input)
		if !ascendingOrder(tt.input) {
			t.Error("order is not valid")
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	b.ReportAllocs()
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})
}
