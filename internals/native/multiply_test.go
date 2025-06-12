package native

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMultiply_Working(t *testing.T) {
	size := 4
	matrixA := NewRandomMatrix(size)
	matrixB := NewRandomMatrix(size)

	result, err := matrixA.Multiply(matrixB)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, size, len(*result))
	require.Equal(t, size, len((*result)[0]))
	require.Equal(t, len(*matrixA), len(*result))           // rows should be the same length as the first matrix
	require.Equal(t, len((*matrixB)[0]), len((*result)[0])) // columns should be the same length as the second matrix

	verify := VerifyMatrixMultiplication(matrixA, matrixB, result)
	require.True(t, verify)
}

func TestMultiply_NotSameLength(t *testing.T) {
	sizeA := 4
	sizeB := 3
	matrixA := NewRandomMatrix(sizeA)
	matrixB := NewRandomMatrix(sizeB)

	result, err := matrixA.Multiply(matrixB)
	require.Error(t, err)
	require.Empty(t, result)
}
