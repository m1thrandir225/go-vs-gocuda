package native

import (
	"fmt"
	"github.com/m1thrandir225/go-vs-gocuda/util"
	"sync"
	"time"
)

func (m *Matrix) MultiplyParallel(b *Matrix) (*Matrix, error) {
	defer util.TimeTrack(time.Now(), "native-parallel")
	if len(*m) == 0 || len((*m)[0]) == 0 || len(*b) == 0 || len((*b)[0]) == 0 {
		return nil, fmt.Errorf("input matricies cannot be empty")
	}
	if len((*m)[0]) != len(*b) {
		return nil, fmt.Errorf("incompattible matrix dimensions for multiplication")
	}

	result := make([][]float64, len(*m))
	for i := range result {
		result[i] = make([]float64, len((*b)[0]))
	}

	var wg sync.WaitGroup

	for i := range *m {
		wg.Add(1)
		go func(rowIdx int) {
			defer wg.Done()
			for j := range (*b)[0] {
				sum := 0.0
				for k := range *b {
					sum += (*m)[rowIdx][k] * (*b)[k][j]
				}
				result[rowIdx][j] = sum
			}
		}(i)
	}

	wg.Wait()

	resultMatrix := Matrix(result)
	return &resultMatrix, nil
}
