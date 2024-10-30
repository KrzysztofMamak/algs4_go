package main

func Lg(n int) int {
	log := 0

	for n > 0 {
		log++
		n /= 2
	}

	return log - 1
}
