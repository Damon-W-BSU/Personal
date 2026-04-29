package sorting

import "os"

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

// dumps bytes to a text file for viewing
func DumpToTxt(raw []byte) {
	f, err := os.Create("sorting/dump.txt")
	if err == nil {
		f.Write(raw)
	}
}
