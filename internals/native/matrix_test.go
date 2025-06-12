package native

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewRandomMatrix(t *testing.T) {
	size := 5
	matrix := NewRandomMatrix(size)

	require.NotEmpty(t, matrix)
	require.Equal(t, len(*matrix), size)
	require.Equal(t, len((*matrix)[0]), size)
}

func TestMatrixDimensions(t *testing.T) {
	size := 4
	matrix := NewRandomMatrix(size)

	x, y := matrix.Dimensions()

	require.NotEmpty(t, x)
	require.NotEmpty(t, y)
	require.Equal(t, x, y)
	require.Equal(t, x, size)
	require.Equal(t, y, size)
}
