package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	selectionSortDvunapravleny(ar)

	fmt.Println(ar)
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
