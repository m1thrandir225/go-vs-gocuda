package main

import (
	"flag"
	"log"

	"github.com/m1thrandir225/go-vs-gocuda/internals/cuda"
	"github.com/m1thrandir225/go-vs-gocuda/internals/native"
)

func main() {
	matrixSize := flag.Int("size", 512, "matrix size")
	flag.Parse()

	a := make([][]float64, *matrixSize)
	b := make([][]float64, *matrixSize)
	for i := 0; i < *matrixSize; i++ {
		a[i] = make([]float64, *matrixSize)
		b[i] = make([]float64, *matrixSize)
		for j := 0; j < *matrixSize; j++ {
			a[i][j] = float64(i + j)
			b[i][j] = float64(i + j)
		}
	}

	c, err := cuda.Multiply(a, b)
	if err != nil {
		log.Fatal(err)
		return
	}

	matrixA := native.Matrix(a)
	matrixB := native.Matrix(b)
	matrixC := native.Matrix(c)

	verificationBasic := native.VerifyMatrixMultiplicationFreivalds(&matrixA, &matrixB, &matrixC, 20)
	log.Printf("Verification CUDA Multiply: %t\n", verificationBasic)
	log.Println("--------------------------------")
	c_tiled, err := cuda.MultiplyTiled(a, b)
	if err != nil {
		log.Fatal(err)
		return
	}

	matrixC_tiled := native.Matrix(c_tiled)
	verificationTiled := native.VerifyMatrixMultiplicationFreivalds(&matrixA, &matrixB, &matrixC_tiled, 20)
	log.Printf("Verification CUDA Multiply Tiled: %t\n", verificationTiled)

}
