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

	a := make([][]float32, *matrixSize)
	b := make([][]float32, *matrixSize)
	for i := 0; i < *matrixSize; i++ {
		a[i] = make([]float32, *matrixSize)
		b[i] = make([]float32, *matrixSize)
		for j := 0; j < *matrixSize; j++ {
			a[i][j] = float32(i + j)
			b[i][j] = float32(i + j)
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

	verificationBasic := native.VerifyMatrixMultiplication(&matrixA, &matrixB, &matrixC)
	log.Printf("Verification CUDA Multiply: %t\n", verificationBasic)
	log.Println("--------------------------------")
	c_tiled, err := cuda.MultiplyTiled(a, b)
	if err != nil {
		log.Fatal(err)
		return
	}

	matrixC_tiled := native.Matrix(c_tiled)
	verificationTiled := native.VerifyMatrixMultiplication(&matrixA, &matrixB, &matrixC_tiled)
	log.Printf("Verification CUDA Multiply Tiled: %t\n", verificationTiled)

}
