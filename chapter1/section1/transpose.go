package main

func Transpose(slice [][]int) [][]int {
	m := len(slice[0])
	n := len(slice)
	result := make([][]int, m)

	for i := range result {
		result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = slice[j][i]
		}
	}

	return result
}
