package native

import (
	"fmt"
	"github.com/m1thrandir225/go-vs-gocuda/util"
	"log"
	"runtime"
	"sync"
	"time"
)

func (m *Matrix) MultiplyParallelWorkerPool(b *Matrix) (*Matrix, error) {
	defer util.TimeTrack(time.Now(), "native-parallel-worker-pool")

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

	numWorkers := runtime.NumCPU()

	log.Printf("Number of CPU's available: %d\n", numWorkers)

	jobs := make(chan int, len(*m))

	var wg sync.WaitGroup

	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for rowIdx := range jobs {
				for j := range (*b)[0] {
					sum := 0.0
					for k := range *b {
						sum += (*m)[rowIdx][k] * (*b)[k][j]
					}
					result[rowIdx][j] = sum
				}
			}
		}()
	}

	for i := range *m {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

	resultMatrix := Matrix(result)

	return &resultMatrix, nil
}
