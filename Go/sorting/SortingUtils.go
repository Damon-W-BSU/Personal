package sorting

// swaps the values of two integers
//
func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// Checks whether an array of integers is sorted
func IsSorted(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}
