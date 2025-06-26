package cuda

import (
	"fmt"
	"log"
	"time"

	"github.com/m1thrandir225/go-vs-gocuda/util"
)

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -L "C:/Program Files/NVIDIA GPU Computing Toolkit/CUDA/v12.9/lib/x64" -lmatrixmult -lcudart
void matrix_multiplication_wrapper(float* a, float* b, float* c, int size);
*/
import "C"

func Multiply(a, b [][]float32) ([][]float32, error) {
	defer util.TimeTrack(time.Now(), "Total Time Multiply CUDA")
	size := len(a)
	if size == 0 || len(a[0]) != size || len(b) != size || len(b[0]) != size {
		return nil, fmt.Errorf("invalid or non-square matrices")
	}

	flatA := flatten(a)
	flatB := flatten(b)
	flatC := make([]float32, size*size)

	beforeCCall := time.Now()

	C.matrix_multiplication_wrapper(
		(*C.float)(&flatA[0]),
		(*C.float)(&flatB[0]),
		(*C.float)(&flatC[0]),
		C.int(size),
	)

	elapsed := time.Since(beforeCCall)

	log.Printf("Go-C CUDA Call Took: %s\n", elapsed)

	result := unflatten(flatC, size)

	return result, nil
}

func flatten(matrix [][]float32) []float32 {
	size := len(matrix)
	flat := make([]float32, 0, size*size)
	for _, row := range matrix {
		flat = append(flat, row...)
	}
	return flat
}

func unflatten(flat []float32, size int) [][]float32 {
	matrix := make([][]float32, size)
	for i := 0; i < size; i++ {
		matrix[i] = flat[i*size : (i+1)*size]
	}
	return matrix
}

func PrintMatrix(matrix [][]float32) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%8.2f", val)
		}
		fmt.Println()
	}
}
