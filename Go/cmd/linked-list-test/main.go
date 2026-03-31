package linkedlist

import (
	"fmt"

	"github.com/Damon-W-BSU/Personal/Go/linkedlist"
)

// basic test suite for list.go
func TestList() {

	L := linkedlist.New[int]()
	fmt.Println(L)

	// append stuff
	fmt.Println("\nAPPEND")
	for i := range 5 {
		L.Append(i)
		fmt.Println(L)
	}

	// insert stuff
	fmt.Println("\nINSERT")
	err := L.Insert(0, 0)
	fmt.Println(L, err)
	err = L.Insert(5, 5)
	fmt.Println(L, err)
	err = L.Insert(3, 3)
	fmt.Println(L, err)
	err = L.Insert(10, 10)
	fmt.Println(L, err)

	// peek stuff
	fmt.Println("\nPEEK")
	L = linkedlist.New[int]()
	for i := range 3 {
		L.Append(i)
	}
	fmt.Println(L)
	for i := range 4 {
		fmt.Println(L.Peek(i))
	}

	// take stuff
	fmt.Println("\nTAKE")
	fmt.Println(L)
	out, err := L.Take(1)
	fmt.Println(L, out, err)
	out, err = L.Take(0)
	fmt.Println(L, out, err)
	out, err = L.Take(1)
	fmt.Println(L, out, err)
	out, err = L.Take(0)
	fmt.Println(L, out, err)

	// indexOf stuff
	fmt.Println("\nINDEX OF")
	for i := range 5 {
		L.Append(i)
	}
	fmt.Println(L)
	fmt.Println(L.IndexOf(0))
	fmt.Println(L.IndexOf(4))
	fmt.Println(L.IndexOf(2))
	fmt.Println(L.IndexOf(5))

	// delete stuff
	fmt.Println("\nDELETE")
	fmt.Println(L)
	out = L.Delete(0)
	fmt.Println(L, out)
	out = L.Delete(4)
	fmt.Println(L, out)
	out = L.Delete(2)
	fmt.Println(L, out)
	out = L.Delete(5)
	fmt.Println(L, out)

}
