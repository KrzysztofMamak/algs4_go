package main

func Histogram(a []int, m int) []int {
	result := make([]int, m)

	for _, v := range a {
		if v < m {
			result[v]++
		}
	}

	return result
}
