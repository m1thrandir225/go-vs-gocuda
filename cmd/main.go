package main

import (
	"fmt"

	"github.com/m1thrandir225/go-vs-gocuda/internals/native"
	"github.com/m1thrandir225/go-vs-gocuda/util"
)

func main() {
	matrixSize := 512
	matrixA := util.CreateMatrix(matrixSize)
	matrixB := util.CreateMatrix(matrixSize)

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
