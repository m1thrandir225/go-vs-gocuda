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
void tiled_matrix_multiplication_wrapper(double  *a, double *b, double *c, int size);
*/
import "C"

func MultiplyTiled(a, b [][]float64) ([][]float64, error) {
	defer util.TimeTrack(time.Now(), "Multiply Tiled CUDA")

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

	log.Printf("Multiply Tiled Go-C Call Took: %s\n", elapsed)

	result := unflatten(flatC, size)

	return result, nil
}
