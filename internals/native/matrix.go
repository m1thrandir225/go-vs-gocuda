package native

import (
	"fmt"
	"github.com/m1thrandir225/go-vs-gocuda/util"
	"math"
)

type Matrix [][]float64

func NewMatrix(size int) *Matrix {
	created := util.CreateMatrix(size)

	matrix := Matrix(created)

	return &matrix
}

func (m *Matrix) Dimensions() (int, int) {
	return len(*m), len(*m)
}

func (m *Matrix) Print() {
	for _, row := range *m {
		for _, col := range row {
			fmt.Printf("%f", col)
		}
		fmt.Println()
	}
}

func (m *Matrix) VerifyMultiplication(a, b, expected *Matrix) bool {
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
	for i := 0; i < rowsA; i++ {
		calculatedResultData[i] = make([]float64, colsA)
		for j := 0; j < colsB; j++ {
			var sum float64
			for k := 0; k < colsA; k++ {
				sum += (*a)[i][k] * (*b)[k][j]
			}
			calculatedResultData[i][j] = sum
		}
	}

	calculatedMatrix := Matrix(calculatedResultData)

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsA; j++ {
			if math.Abs(calculatedMatrix[i][j]-(*expected)[i][j]) > epsilon {
				fmt.Printf("Error: matrix values do not match.")
				return false
			}
		}
	}

	return true
}
