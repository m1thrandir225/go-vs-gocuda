package native

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMatrix_MultiplyParallel_Working(t *testing.T) {
	size := 4
	matrixA := NewRandomMatrix(size)
	matrixB := NewRandomMatrix(size)

	result, err := matrixA.MultiplyParallel(matrixB)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, size, len(*result))
	require.Equal(t, size, len((*result)[0]))
	require.Equal(t, len(*matrixA), len(*result))
	require.Equal(t, len((*matrixB)[0]), len((*result)[0]))
	verify := VerifyMatrixMultiplication(matrixA, matrixB, result)
	require.True(t, verify)
}

func TestMatrix_MultiplyParallel_DifferentSizes(t *testing.T) {
	sizeA := 4
	sizeB := 3
	matrixA := NewRandomMatrix(sizeA)
	matrixB := NewRandomMatrix(sizeB)
	result, err := matrixA.MultiplyParallel(matrixB)

	require.Error(t, err)
	require.Empty(t, result)
}

func TestMatrix_MultiplyParallel_EmptyMatrix(t *testing.T) {
	size := 2
	matrixA := NewRandomMatrix(size)
	matrixBase := make([][]float64, size)
	matrixB := Matrix(matrixBase)

	result, err := matrixA.MultiplyParallel(&matrixB)
	require.Error(t, err)
	require.Empty(t, result)
}
