package cuda

import (
	"fmt"
	"log"
	"time"
	"unsafe"

	"github.com/m1thrandir225/go-vs-gocuda/util"
)

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -L "C:/Program Files/NVIDIA GPU Computing Toolkit/CUDA/v12.9/lib/x64" -lmatrixmult -lcudart
#include "matrix_mult.h"
void matrix_multiplication_wrapper(double *a, double *b, double *c, int size);
*/
import "C"

func Multiply(a, b [][]float64) ([][]float64, error) {
	defer util.TimeTrack(time.Now(), "Multiply CUDA")
	size := len(a)
	if size == 0 || len(a[0]) != size || len(b) != size || len(b[0]) != size {
		return nil, fmt.Errorf("invalid or non-square matrices")
	}

	flatA := flatten(a)
	flatB := flatten(b)
	flatC := make([]float64, size*size)

	beforeCCall := time.Now()

	C.matrix_multiplication_wrapper(
		(*C.double)(unsafe.Pointer(&flatA[0])),
		(*C.double)(unsafe.Pointer(&flatB[0])),
		(*C.double)(unsafe.Pointer(&flatC[0])),
		C.int(size),
	)

	elapsed := time.Since(beforeCCall)
	log.Printf("Multiply Go-C Call Took: %s\n", elapsed)

	result := unflatten(flatC, size)

	return result, nil
}

func flatten(matrix [][]float64) []float64 {
	size := len(matrix)
	flat := make([]float64, 0, size*size)
	for _, row := range matrix {
		flat = append(flat, row...)
	}
	return flat
}

func unflatten(flat []float64, size int) [][]float64 {
	matrix := make([][]float64, size)
	for i := 0; i < size; i++ {
		matrix[i] = flat[i*size : (i+1)*size]
	}
	return matrix
}

func PrintMatrix(matrix [][]float64) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%8.2f", val)
		}
		fmt.Println()
	}
}
