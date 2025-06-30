package native

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/m1thrandir225/go-vs-gocuda/util"
)

type Matrix [][]float64

// Creates a random new matrix
func NewRandomMatrix(size int) *Matrix {
	created := util.CreateMatrix(size)

	matrix := Matrix(created)

	return &matrix
}

// Dimensions Returns the dimensions of the current matrix
func (m *Matrix) Dimensions() (int, int) {
	return len(*m), len(*m)
}

// Print Prints out the matrix to standard output
func (m *Matrix) Print() {
	for _, row := range *m {
		for _, col := range row {
			fmt.Printf("%f", col)
		}
		fmt.Println()
	}
}

// VerifyMatrixMultiplication Verifies if a multiplication of two matricies is
// correct
func VerifyMatrixMultiplication(a, b, expected *Matrix) bool {
	epsilon := 1e-9
	if a == nil || b == nil || expected == nil {
		fmt.Printf("Error: one or more matricies are nil.")
		return false
	}

	rowsA, colsA := len(*a), len((*a)[0])
	rowsB, colsB := len(*b), len((*b)[0])
	rowsExpected, colsExpected := len(*expected), len((*expected)[0])

	if colsA != rowsB {
		fmt.Printf("Error: Cannot multiply matrices. Columns of A (%d) != Rows of B (%d).\n", colsA, rowsB)
		return false
	}

	if rowsA != rowsExpected || colsB != colsExpected {
		fmt.Printf("Error: Expected result matrix has wrong dimensions. Expected %dx%d, got %dx%d.\n", rowsA, colsB, rowsExpected, colsExpected)
		return false
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			var sum float64
			for k := 0; k < colsA; k++ {
				sum += (*a)[i][k] * (*b)[k][j]
			}
			if math.Abs(float64(sum-(*expected)[i][j])) > epsilon {
				fmt.Printf("Error: matrix values do not match at (%d, %d). Expected %f, got %f\n", i, j, (*expected)[i][j], sum)
				return false
			}
		}
	}

	return true
}

// Faster Verification Algorithm
func VerifyMatrixMultiplicationFreivalds(a, b, expected *Matrix, iterations int) bool {
	relativeTolerance := 1e-12

	if a == nil || b == nil || expected == nil {
		fmt.Println("Error: one or more matrices are nil.")
		return false
	}

	rowsA, colsA := len(*a), len((*a)[0])
	rowsB, colsB := len(*b), len((*b)[0])
	rowsExpected, colsExpected := len(*expected), len((*expected)[0])

	if colsA != rowsB {
		fmt.Printf("Error: Cannot multiply matrices. Columns of A (%d) != Rows of B (%d).\n", colsA, rowsB)
		return false
	}
	if rowsA != rowsExpected || colsB != colsExpected {
		fmt.Printf("Error: Expected result matrix has wrong dimensions. Expected %dx%d, got %dx%d.\n", rowsA, colsB, rowsExpected, colsExpected)
		return false
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for k := 0; k < iterations; k++ {
		// 1. Generate random vector 'r'
		r := make([]float64, colsB)
		for i := range r {
			r[i] = rng.Float64()
		}

		// 2. Compute Br = B * r
		br := make([]float64, rowsB)
		for i := 0; i < rowsB; i++ {
			var sum float64
			for j := 0; j < colsB; j++ {
				sum += (*b)[i][j] * r[j]
			}
			br[i] = sum
		}

		// 3. Compute ABr = A * (B*r)
		abr := make([]float64, rowsA)
		for i := 0; i < rowsA; i++ {
			var sum float64
			for j := 0; j < colsA; j++ {
				sum += (*a)[i][j] * br[j]
			}
			abr[i] = sum
		}

		// 4. Compute Cr = expected * r
		cr := make([]float64, rowsExpected)
		for i := 0; i < rowsExpected; i++ {
			var sum float64
			for j := 0; j < colsExpected; j++ {
				sum += (*expected)[i][j] * r[j]
			}
			cr[i] = sum
		}

		// 5. Compare ABr and Cr using a relative epsilon
		for i := 0; i < rowsA; i++ {
			diff := math.Abs(abr[i] - cr[i])
			magnitude := math.Max(math.Abs(abr[i]), math.Abs(cr[i]))
			epsilon := 1e-9 + relativeTolerance*magnitude

			if diff > epsilon {
				fmt.Println("--------------------------------------------------")
				fmt.Printf("Error: Matrix verification failed on iteration %d, row %d\n", k, i)
				fmt.Printf("  Computed A*(B*r):  %.15f\n", abr[i])
				fmt.Printf("  Computed C*r:      %.15f\n", cr[i])
				fmt.Printf("  Difference:        %.15f\n", diff)
				fmt.Printf("  Required Epsilon:  %.15f\n", epsilon)
				fmt.Println("--------------------------------------------------")
				return false
			}
		}
	}

	return true
}
