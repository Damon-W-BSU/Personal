package sorting

import (
	"math/rand"
)

// Randomizes the order of an array of integers
func Randomize(list []int) {
	for i := range list {
		swapIndex := (len(list) - 1) - rand.Int()%(len(list)-i)
		Swap(&list[i], &list[swapIndex])
	}
}

// Sorts an array of integers by repeatedly randomzing order until sorted
func RandomSort(list []int) {
	for !IsSorted(list) {
		Randomize(list)
	}
}
