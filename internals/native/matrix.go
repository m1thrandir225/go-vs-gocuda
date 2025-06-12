package native

import (
	"fmt"
	"math"

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

	calculatedResultData := make([][]float64, rowsA)
	for i := range rowsA {
		calculatedResultData[i] = make([]float64, colsA)
		for j := range colsB {
			var sum float64
			for k := range rowsA {
				sum += (*a)[i][k] * (*b)[k][j]
			}
			calculatedResultData[i][j] = sum
		}
	}

	calculatedMatrix := Matrix(calculatedResultData)

	for i := range rowsA {
		for j := range colsA {
			if math.Abs(calculatedMatrix[i][j]-(*expected)[i][j]) > epsilon {
				fmt.Printf("Error: matrix values do not match.")
				return false
			}
		}
	}

	return true
}
