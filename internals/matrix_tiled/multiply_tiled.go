package main

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR} -L"C:/Program Files/NVIDIA GPU Computing Toolkit/CUDA/v12.9/lib/x64" -lmulttiled -lcudart
void matrix_multiplication_tiled_wrapper(float* a, float* b, float* c, int size);
*/
import "C"
import (
	"fmt"
)

func Multiply(a, b [][]float32) ([][]float32, error) {
	// Get matrix dimensions (assuming square matrices)
	size := len(a)
	if size == 0 || len(a[0]) != size || len(b) != size || len(b[0]) != size {
		return nil, fmt.Errorf("invalid or non-square matrices")
	}

	// Flatten the 2D Go slices into 1D slices, which is how C expects arrays.
	flatA := flatten(a)
	flatB := flatten(b)
	flatC := make([]float32, size*size)

	// Call the C function from the shared library.
	// We pass pointers to the first element of our Go slices.
	// unsafe.Pointer is Go's way of working with raw C pointers.
	C.matrix_multiplication_tiled_wrapper(
		(*C.float)(&flatA[0]),
		(*C.float)(&flatB[0]),
		(*C.float)(&flatC[0]),
		C.int(size),
	)

	// "Un-flatten" the 1D result slice back into a 2D Go slice.
	result := unflatten(flatC, size)

	return result, nil
}

// flatten converts a 2D slice of float32 to a 1D slice.
func flatten(matrix [][]float32) []float32 {
	size := len(matrix)
	flat := make([]float32, 0, size*size)
	for _, row := range matrix {
		flat = append(flat, row...)
	}
	return flat
}

// unflatten converts a 1D slice back to a 2D slice.
func unflatten(flat []float32, size int) [][]float32 {
	matrix := make([][]float32, size)
	for i := 0; i < size; i++ {
		matrix[i] = flat[i*size : (i+1)*size]
	}
	return matrix
}

func printMatrix(matrix [][]float32) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%8.2f", val)
		}
		fmt.Println()
	}
}

func main() {
	const N = 4
	a := make([][]float32, N)
	b := make([][]float32, N)

	for i := 0; i < N; i++ {
		a[i] = make([]float32, N)
		b[i] = make([]float32, N)
		for j := 0; j < N; j++ {
			a[i][j] = float32(i*N + j) // Simple values
			b[i][j] = float32(j*N + i) // Transposed values
		}
	}

	fmt.Println("Matrix A:")
	printMatrix(a)
	fmt.Println("\nMatrix B:")
	printMatrix(b)
	// Perform multiplication
	c, err := Multiply(a, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nResult Matrix C (from CUDA):")
	printMatrix(c)
}
