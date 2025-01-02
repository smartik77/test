package sort

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func bubbleSort(ar []int) {
	isSorted := false
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar)-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j+1], ar[j] = ar[j], ar[j+1]
				isSorted = true
			}
		}
		if !isSorted {
			return
		}
	}
}

func bubbleSortReversed(ar []int) {
	isSorted := false
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar)-1; j++ {
			if ar[j] < ar[j+1] {
				ar[j+1], ar[j] = ar[j], ar[j+1]
				isSorted = true
			}
		}
		if !isSorted {
			return
		}
	}
}

func selectionSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			ar[i], ar[minIndex] = ar[minIndex], ar[i]
		}
	}
}

func selectionSortByMax(ar []int) {
	for i := len(ar) - 1; i > 0; i-- {
		maxIndex := i
		for j := i - 1; j >= 0; j-- {
			if ar[j] > ar[maxIndex] {
				maxIndex = j
			}
		}
		if maxIndex != i {
			ar[i], ar[maxIndex] = ar[maxIndex], ar[i]
		}
	}
}

func selectionSortDvunapravleny(ar []int) {
	for i := 0; i < len(ar)/2; i++ {
		minIndex := i
		maxIndex := i
		for j := i + 1; j < len(ar)-i; j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
			if ar[j] > ar[maxIndex] {
				maxIndex = j
			}
		}
		if minIndex != i {
			ar[i], ar[minIndex] = ar[minIndex], ar[i]
			if maxIndex == i {
				maxIndex = minIndex
			}
		}
		if maxIndex != len(ar)-i-1 {
			ar[len(ar)-i-1], ar[maxIndex] = ar[maxIndex], ar[len(ar)-i-1]
		}
	}
}

func insertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		key := ar[i]
		j := i - 1
		for j >= 0 && ar[j] > key {
			ar[j+1] = ar[j]
			j--
		}
		ar[j+1] = key
	}
}

func mergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	middle := len(ar) / 2

	sortedAr := make([]int, 0, len(ar))
	left, right := mergeSort(ar[:middle]), mergeSort(ar[middle:])

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			sortedAr = append(sortedAr, right[j])
			j++
		} else {
			sortedAr = append(sortedAr, left[i])
			i++
		}
	}

	sortedAr = append(sortedAr, left[i:]...)
	sortedAr = append(sortedAr, right[j:]...)

	return sortedAr
}

func quickSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	left, right := 0, len(ar)-1
	pivotIndex := rand.Int() % len(ar)

	ar[pivotIndex], ar[right] = ar[right], ar[pivotIndex]

	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}

	ar[left], ar[right] = ar[right], ar[left]

	quickSort(ar[:left])
	quickSort(ar[left+1:])

	return
}
