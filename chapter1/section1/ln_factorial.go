package main

import "math"

// https://tutorial.math.lamar.edu/classes/alg/logfunctions.aspx#MathJax-Element-87-Frame
// ln(N!) = ln(N * (N - 1)!) = ln(N) + ln((N - 1)!)
// logb(xy) = logb(x) + logb(y)
func LnFactorial(n int) float64 {
	if n == 1 {
		return 0
	}

	return math.Log(float64(n)) + LnFactorial(n -1)
}
