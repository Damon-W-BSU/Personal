package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello, World!")

	/*
		VARIABLES
	*/

	fmt.Print("\nVARIABLES\n\n")

	// "var" declares list of variables
	var a, b int
	a = 0
	b = 1
	fmt.Println(a, b)

	// when initialized, no type is needed
	var c, d = 5, "Hello!"
	fmt.Println(c, d)

	// type can be implicit with ":=" in the declaration
	// this cannot be used outside of a function
	e := 3.14
	fmt.Println(e)

	// type conversions are done with T(v)
	f := int(e)
	fmt.Println(f)

	/*
		CONTROL FLOW
	*/

	fmt.Print("\nCONTROL FLOW\n\n")

	// Go only has for loops comprised of init, condition, and post components
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// init and post statements are optional (this is essentially a while loop)
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// if statements can also have init components that remain in scope
	if x := rand.Int() % 10; x < 5 {
		fmt.Println("x is less than 5!")
	} else {
		fmt.Println("x is not less than 5!")
	}

	// switch statements don't need break statements
	// they also need not be integers or constants
	x := rand.Int() % 2
	switch x {
	case 0:
		fmt.Println("x is zero")
	case 1:
		fmt.Println("x is one")
	default:
		fmt.Println("x is two")
	}

	// can just say "switch" for the equivelant of "switch true" for if-else chains
	x = rand.Int() % 10
	switch {
	case x <= 3:
		fmt.Println("x is 3 or less")
	case x <= 7:
		fmt.Println("x is between 4-7")
	default:
		fmt.Println("x is greater than 7")
	}

	// defer statements defer execution of statement until function call completes
	// defer statements are pushed to a stack if multiple are used
	// side note: this is an anonymous function
	func() {
		defer fmt.Println("first")
		defer fmt.Println("second")
		fmt.Println("third")
	}()

	/*
		POINTERS
	*/

	fmt.Print("\nPOINTERS\n\n")

	// pointers are declared using an asterix and the type to point to
	var pointer *int
	pointerValue := 5

	// and an ampersand to assign operand to point to
	pointer = &pointerValue

	// pointers are dereferenced with an asterix as well
	fmt.Println(pointerValue, pointer, *pointer)

	/*
		STRUCTS
	*/

	fmt.Print("\nSTRUCTS\n\n")

	// structs are simply a collection of fields
	type Cat struct {
		name string
		age  int
	}

	myCat := Cat{"Ace", 5}

	// fields are accesed using dot notation
	fmt.Println(myCat.name, myCat.age)

	// this extends to pointer notation
	p := &myCat
	fmt.Println(p.name, p.age)

	/*
		ARRAYS
	*/

	fmt.Print("\nARRAYS\n\n")

	// Arrays are declared with [n]T where n is the length
	// Array length is immutable, but this can be solved with slices
	var list [2]int
	list[0] = 10
	list[1] = 20
	fmt.Println(list)

	// To declare and initialize simultaneously, use brackets
	names := [2]string{"Ace", "Chuck"}
	fmt.Println(names)

	/*
		SLICES
	*/

	fmt.Print("\nSLICES\n\n")

	// slices are references to underlying arrays; changing a slice changes the array
	// slices are created by specifying low (included) and high(exluded) indices of an array
	var nums = [5]int{1, 2, 3, 4, 5}
	var numSlice []int = nums[1:4]
	fmt.Println(numSlice)

	// slices literals can also be initilized and an underlying array will also be allocated
	soloSlice := []int{6, 7, 8, 9, 10}
	fmt.Println(soloSlice)

	// all slices have length and capacity values
	s := soloSlice[:3]
	fmt.Println("New slice length:", len(s))
	fmt.Println("New slice capacity:", cap(s))

	// slices can also be allocated using make
	s = make([]int, 5)
	fmt.Println(s)

	// multi-dimension slices
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(grid)

	// the append function allows elements to be added and automatically reallocated
	s = []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println(s)

	// the "range" keyword in a loop iterates over both indices and items
	for index, value := range s {
		fmt.Printf("index: %d, value: %d", index, value)
	}

	

}

/*
	FUNCTIONS
*/

// basic function header
// Note: capitilization exports function
func Add(x int, y int) int {
	return x + y
}

// function headers can name the return type
func Subtract(x int, y int) (z int) {
	z = x + y
	return
}

// function headers with parameters of the same type can be shortened
func Multiply(x, y int) int {
	return x * y
}

// functions can also have multiple return values
func Divide(x, y int) (int, int) {
	var z int = x / y
	var remainder int = x % y
	return z, remainder
}
