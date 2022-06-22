package main

func isPrime(n int) bool {
	for x := 2; x < n; x++ {
		if n%x == 0 {
			return false
		}
	}
	return true
}

func FindPrimeByRange(start, end int) []int {
	var prime []int

	for i := start; i <= end; i++ {
		if isPrime(i) {
			prime = append(prime, i)
		}
	}

	return prime
}
