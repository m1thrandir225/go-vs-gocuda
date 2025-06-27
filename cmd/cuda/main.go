package main

import (
	"fmt"

	"github.com/m1thrandir225/go-vs-gocuda/internals/cuda"
	"github.com/m1thrandir225/go-vs-gocuda/internals/native"
)

func main() {
	const N = 2048
	a := make([][]float32, N)
	b := make([][]float32, N)
	for i := 0; i < N; i++ {
		a[i] = make([]float32, N)
		b[i] = make([]float32, N)
		for j := 0; j < N; j++ {
			a[i][j] = float32(i + j)
			b[i][j] = float32(i + j)
		}
	}
	// Perform multiplication
	c, err := cuda.Multiply(a, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	matrixA := native.Matrix(a)
	matrixB := native.Matrix(b)
	matrixC := native.Matrix(c)

	verfied := native.VerifyMatrixMultiplication(&matrixA, &matrixB, &matrixC)

	if !verfied {
		fmt.Println("Error: Verification failed")
		return
	}
	fmt.Println("Verification: True")

	c_tiled, err := cuda.MultiplyTiled(a, b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	matrixC_tiled := native.Matrix(c_tiled)

	verify_tiled := native.VerifyMatrixMultiplication(&matrixA, &matrixB, &matrixC_tiled)

	if !verify_tiled {
		fmt.Println("Error: Tiled Matrix Multiplication Verification Failed")
		return
	}

	fmt.Println("Tiled Matrix Mult. Verifiaction: True")
}
