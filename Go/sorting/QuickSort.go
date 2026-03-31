package sorting

import (
	"math/rand"
)

// Sorts a slice of integers, ascending
func QuickSort(list []int) {

	// base case
	if len(list) <= 1 {
		return
	}

	// partition and recursivrly sort resulting slices
	left, right := partition(list)
	QuickSort(left)
	QuickSort(right)
}

// Splits a slice of integers into two slices containing the
// integers above and below a randomly selected pivot, respectively
func partition(list []int) ([]int, []int) {

	// get random index and move to front
	index := rand.Int() % len(list)
	list[index], list[0] = list[0], list[index]
	pivot := list[0]

	// arrange list items based on pivot
	swap := len(list) - 1 // swapping index
	for i := 1; i <= swap; {
		if list[i] > pivot {
			list[i], list[swap] = list[swap], list[i]
			swap--
		} else {
			i++
		}
	}

	// move pivot into place
	list[0], list[swap] = list[swap], list[0]

	return list[:swap], list[swap+1:]
}
