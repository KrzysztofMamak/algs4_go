package main

func Binominal(n int, k int, p float64) float64 {
	// Base case
	if n == 0 && k == 0 {
		return 1
	}

	// Invalid input
	if n < 0 || k < 0 {
		return 0
	}

	// P(A or B) = P(A) + P(B)
	return (1-p)*Binominal(n-1, k, p) + p*Binominal(n-1, k-1, p)
}
