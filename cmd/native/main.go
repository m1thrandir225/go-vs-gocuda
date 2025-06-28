package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/m1thrandir225/go-vs-gocuda/internals/native"
)

func main() {
	matrixSize := flag.Int("size", 512, "matrix size")
	flag.Parse()

	matrixA := native.NewRandomMatrix(*matrixSize)
	matrixB := native.NewRandomMatrix(*matrixSize)

	resultNative, err := matrixA.Multiply(matrixB)
	if err != nil {
		log.Fatal(err)
		return
	}
	verificationNative := native.VerifyMatrixMultiplication(matrixA, matrixB, resultNative)
	log.Printf("Verification native multiplication: %t\n", verificationNative)

	resultParallel, err := matrixA.MultiplyParallel(matrixB)
	if err != nil {
		log.Fatal(err)
		return
	}

	verificationParallel := native.VerifyMatrixMultiplication(matrixA, matrixB, resultParallel)
	log.Printf("Verification parallel multiplication: %t\n", verificationParallel)

	resultParallelWorkerPool, err := matrixA.MultiplyParallelWorkerPool(matrixB)
	if err != nil {
		log.Fatal(err)
		return
	}

	verificationParallelWorkerPool := native.VerifyMatrixMultiplication(matrixA, matrixB, resultParallelWorkerPool)
	fmt.Printf("Verification worker pool multiplication: %t\n", verificationParallelWorkerPool)
}
