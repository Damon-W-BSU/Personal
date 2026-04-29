package syntax

// generates the first n prime numbers and returns slice containing primes
func GeneratePrimes(amount int) []int {

	primes := make([]int, 0, amount)
	current := 2

	for len(primes) < amount {
		isPrime := true
		for i := 2; i*i <= current; i++ {
			if current%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, current)

		}
		current++
	}

	return primes
}
