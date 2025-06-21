package main

import (
	"fmt"
	"github.com/m1thrandir225/go-vs-gocuda/internals/cuda"
)

func main() {
	const N = 4

	a := make([][]float32, N)
	b := make([][]float32, N)
	for i := 0; i < N; i++ {
		a[i] = make([]float32, N)
		b[i] = make([]float32, N)
		for j := 0; j < N; j++ {
			a[i][j] = float32(i*N + j)
			b[i][j] = float32(j*N + i)
		}
	}

	fmt.Println("Matrix A:")
	cuda.PrintMatrix(a)
	fmt.Println("\nMatrix B:")
	cuda.PrintMatrix(b)

	// Perform multiplication
	c, err := cuda.Multiply(a, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nResult Matrix C (from CUDA):")
	cuda.PrintMatrix(c)
}
