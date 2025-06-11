package main

import (
	"fmt"
	"github.com/m1thrandir225/go-vs-gocuda/util"

	"github.com/m1thrandir225/go-vs-gocuda/internals/native"
)

func main() {
	matrixSize := 512
	matrixA := native.NewRandomMatrix(matrixSize)
	matrixB := native.NewRandomMatrix(matrixSize)

	resultNative, err := matrixA.Multiply(matrixB)
	if err != nil {
		fmt.Println(err)
		return
	}
	verificationNative := util.VerifyMatrixMultiplication(matrixA, matrixB, resultNative)
	fmt.Printf("Verification native multiplication: %t\n", verificationNative)

	resultParallel, err := matrixA.MultiplyParallel(matrixB)
	if err != nil {
		fmt.Println(err)
		return
	}

	verificationParallel := util.VerifyMatrixMultiplication(matrixA, matrixB, resultParallel)
	fmt.Printf("Verification parallel multiplication: %t\n", verificationParallel)

	resultParallelWorkerPool, err := matrixA.MultiplyParallelWorkerPool(matrixB)
	if err != nil {
		fmt.Println(err)
		return
	}

	verificationParallelWorkerPool := util.VerifyMatrixMultiplication(matrixA, matrixB, resultParallelWorkerPool)
	fmt.Printf("Verification worker pool multiplication: %t\n", verificationParallelWorkerPool)
}
