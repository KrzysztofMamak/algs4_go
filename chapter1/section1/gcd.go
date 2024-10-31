package main

func Gcd(p int, q int) int {
	if q == 0 {
		return p
	} else {
		return Gcd(q, p%q)
	}
}

// https://cp-algorithms.com/algebra/extended-euclid-algorithm.html
func ExtendedGcd(a int, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}

	gcd, x1, y1 := ExtendedGcd(b, a%b)
	
	return gcd, y1, x1 - (a/b)*y1
}
