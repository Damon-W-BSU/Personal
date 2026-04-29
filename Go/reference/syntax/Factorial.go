package syntax

func Factorial_iter(num int) int {

	result := 1

	for i := num; i > 1; i-- {
		result *= i
	}

	return result
}

func Factorial_rec(num int) int {
	if num <= 1 {
		return 1
	}
	return num * Factorial_rec(num-1)
}
