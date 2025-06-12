package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateMatrix(t *testing.T) {
	size := 5
	matrix := CreateMatrix(size)

	require.NotEmpty(t, matrix)
	require.Equal(t, size, len(matrix))
	require.Equal(t, size, len(matrix[0]))

	for _, v := range matrix {
		require.NotEmpty(t, v)
	}
}
