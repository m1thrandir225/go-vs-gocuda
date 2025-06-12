package native

import (
	"bytes"
	"io"
	"os"
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

func TestMatrix_Print(t *testing.T) {
	matrix := NewRandomMatrix(1)

	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	matrix.Print()
	w.Close()

	os.Stdout = originalStdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	expectedOutput := "0.000000\n"
	require.Equal(t, expectedOutput, buf.String())
}
