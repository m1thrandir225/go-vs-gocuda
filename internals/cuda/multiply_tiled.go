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
#include "matrix_mult.h"
void tiled_matrix_multiplication_wrapper(float* a, float* b, float* c, int size);
*/
import "C"

func MultiplyTiled(a, b [][]float32) ([][]float32, error) {
	defer util.TimeTrack(time.Now(), "Total Time Multiply Tiled CUDA")

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
