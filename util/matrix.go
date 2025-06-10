package util

func CreateMatrix(size int) [][]float64 {
	matrix := make([][]float64, size)
	for i := range size {
		matrix[i] = make([]float64, size)
		for j := range size {
			matrix[i][j] = float64(i + j) // Just some dummy data
		}
	}
	return matrix
}

func CreateMatrix32(size int) [][]float32 {
	matrix := make([][]float32, size)
	for i := range size {
		matrix[i] = make([]float32, size)
		for j := range size {
			matrix[i][j] = float32(i + j) // Just some dummy data
		}
	}
	return matrix
}
