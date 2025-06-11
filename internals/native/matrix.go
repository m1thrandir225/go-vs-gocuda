package native

import (
	"fmt"
	"github.com/m1thrandir225/go-vs-gocuda/util"
	"math"
)

type Matrix [][]float64

// Creates a random new matrix
func NewRandomMatrix(size int) *Matrix {
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
