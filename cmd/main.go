package main

import (
	"fmt"

	"github.com/m1thrandir225/go-vs-gocuda/internals/native"
)

func createMatrix(size int) [][]float64 {
	matrix := make([][]float64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			matrix[i][j] = float64(i + j) // Just some dummy data
		}
	}
	return matrix
}

func main() {
	matrixSize := 2048
	matrixA := createMatrix(matrixSize)
	matrixB := createMatrix(matrixSize)

	_, err := native.Multiply(matrixA, matrixB)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = native.MultiplyParallel(matrixA, matrixB)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = native.MultiplyParallelWorkerPool(matrixA, matrixB)
	if err != nil {
		fmt.Println(err)
		return
	}
}
