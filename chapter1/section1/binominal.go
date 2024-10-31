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

func BinominalEnhanced(n int, k int, p float64) float64 {
	memo := make([][]float64, n+1)

	for i := range memo {
		memo[i] = make([]float64, k+1)

		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return binominalEnhancedInner(n, k, p, memo)
}

func binominalEnhancedInner(n int, k int, p float64, memo [][]float64) float64 {
	if n == 0 && k == 0 {
		return 1
	}

	if n < 0 || k < 0 {
		return 0
	}

	if memo[n][k] != -1 {
		return memo[n][k]
	}

	memo[n][k] = (1-p)*binominalEnhancedInner(n-1, k, p, memo) + p*binominalEnhancedInner(n-1, k-1, p, memo)

	return memo[n][k]
}
