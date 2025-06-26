package util

func CreateMatrix(size int) [][]float32 {
	matrix := make([][]float32, size)
	for i := range size {
		matrix[i] = make([]float32, size)
		for j := range size {
			matrix[i][j] = float32(i + j)
		}
	}
	return matrix
}
