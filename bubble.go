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

	bubbleSort(ar)

	fmt.Println(ar)

	bubbleSortReversed(ar)

	fmt.Println(ar)
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
