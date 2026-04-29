package sorting

import "sync"

// Uses recursive calls and goroutines to improve mergesort runtime
func CCMergeSort(list []int) []int {

	if len(list) < 2 {
		return list
	}

	left, right := split(list)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		left = MergeSort(left)
	}()
	go func() {
		defer wg.Done()
		right = MergeSort(right)
	}()
	wg.Wait()
	return merge(left, right)
}

// recursively sorts a list of integers in O(nlog(n))
func MergeSort(list []int) []int {

	// base case
	if len(list) < 2 {
		return list
	}

	// split until len 1 reached, then merge
	left, right := split(list)
	left = MergeSort(left)
	right = MergeSort(right)
	return merge(left, right)

}

// splits a list in half
func split(list []int) ([]int, []int) {

	i := len(list) / 2

	left := list[:i]
	right := list[i:]

	return left, right

}

// merges two sorted lists into one sorted list
func merge(left []int, right []int) []int {

	out := []int{}
	var i, j int

	// append lesser value of two lists until one list empty
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			out = append(out, left[i])
			i++
		} else {
			out = append(out, right[j])
			j++
		}
	}

	// append rest of left list
	for i < len(left) {
		out = append(out, left[i])
		i++
	}

	// append rest of right list
	for j < len(right) {
		out = append(out, right[j])
		j++
	}

	return out

}
