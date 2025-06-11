package native

import (
	"fmt"
	"sync"
	"time"

	"github.com/m1thrandir225/go-vs-gocuda/util"
)

func (m *Matrix) Multiply(b *Matrix) (*Matrix, error) {
	defer util.TimeTrack(time.Now(), "native")

	if len((*m)[0]) != len(*b) {
		return nil, fmt.Errorf("incompattible matrix dimensions for multiplication")
	}

	result := make([][]float64, len(*m))
	for i := range result {
		result[i] = make([]float64, len((*b)[0]))
	}

	var wg sync.WaitGroup

	for i := range *m {
		for j := range (*b)[0] {
			for k := range *b {
				result[i][j] += (*m)[i][k] * (*b)[k][j]
			}
		}
	}

	wg.Wait()

	resultMatrix := Matrix(result)

	return &resultMatrix, nil
}
