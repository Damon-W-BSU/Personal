package mathshit

import (
	"fmt"
	"math"
)

// Recursively converts a matrix to row echelon form
func REF(matrix [][]float64) [][]float64 {

	PrintMatrix(matrix)

	// base case
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	innerLen := len(matrix) - 1    // tracks length of inner matrix
	innerWid := len(matrix[0]) - 1 // track width of inner matrix

	for row := range matrix {

		// move to next row if first number in column is 0
		if matrix[row][0] == 0 {
			continue
		}

		// swap rows if not at first row
		if row != 0 {
			temp := matrix[0]
			matrix[0] = matrix[row]
			matrix[row] = temp
		}
		break
	}

	// if zero, only shrink width of inner matrix
	if (math.Abs(0 - matrix[0][0])) < 0.0001 {
		innerLen++
	} else { // perform matrix row operations
		if matrix[0][0] != 1 { // if not already 1, multiply by constant
			div := matrix[0][0]
			for colIndex := range matrix[0] {
				matrix[0][colIndex] /= div
			}
		}

		// Add multiples of row[0] to other rows such that all following
		// values in the matrix's first column equal 0
		var multiplier float64
		for row := 1; row < len(matrix); row++ {
			if matrix[row][0] != 0 {
				multiplier = matrix[row][0] * -1
				for col := range matrix[row] {
					matrix[row][col] = matrix[row][col] + (multiplier * matrix[0][col])
				}
			}
		}

	}

	lenOffset := (len(matrix)) - innerLen
	widOffset := (len(matrix[0])) - innerWid

	// allocate space for matrix containing inner elements
	inner := make([][]float64, innerLen)
	for col := range inner {
		inner[col] = make([]float64, innerWid)
	}

	// populate inner matrix
	for row := 0; row < innerLen; row++ {
		for col := 0; col < innerWid; col++ {
			inner[row][col] = matrix[row+lenOffset][col+widOffset]
		}
	}

	// recursively check remaining matrix
	inner = REF(inner)

	// repopulate and return matrix
	for row := 0; row < len(matrix)-lenOffset; row++ {
		for col := 0; col < len(matrix[0])-widOffset; col++ {
			matrix[row+lenOffset][col+widOffset] = inner[row][col]
		}
	}

	return matrix

}

// Converts a matrix to reduced row echelon form
func RREF(matrix [][]float64) [][]float64 {

	// convert to row echelon form
	matrix = REF(matrix)

	// row major traversal of matrix starting at bottom left
	for row := len(matrix) - 1; row >= 0; row-- {
		for col := 0; col < len(matrix[0]); col++ {

			// If pivot found, add multiples of pivot row such that every other value in the column equals 0
			if matrix[row][col] == 1 {

				for redRow := row - 1; redRow >= 0; redRow-- {

					// FIXME need to edit this condition maybe
					// check for any non zero value in column
					if math.Abs(0-matrix[redRow][col]) > 0.0001 {
						multiplier := matrix[redRow][col] * -1

						// traverse up column, add multiples to row
						for redCol := col; redCol < len(matrix[0]); redCol++ {
							matrix[redRow][redCol] = matrix[redRow][redCol] + (multiplier * matrix[row][redCol])
						}
					}
				}
				break // go to next row
			}
		}
	}
	return matrix
}

// Prints a matrix in grid format
func PrintMatrix(matrix [][]float64) {
	for row := range matrix {
		for col := range matrix[row] {
			fmt.Printf("%10.3f ", matrix[row][col])
		}
		fmt.Println()
	}
	fmt.Println()
}
