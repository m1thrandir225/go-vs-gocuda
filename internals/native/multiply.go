package native

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/m1thrandir225/go-vs-gocuda/util"
)

func Multiply(a, b [][]float64) ([][]float64, error) {
	defer util.TimeTrack(time.Now(), "native")

	if len(a[0]) != len(b) {
		return nil, fmt.Errorf("incompattible matrix dimensions for multiplication")
	}

	result := make([][]float64, len(a))
	for i := range result {
		result[i] = make([]float64, len(b[0]))
	}

	var wg sync.WaitGroup

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	wg.Wait()

	return result, nil
}

func MultiplyParallel(a, b [][]float64) ([][]float64, error) {
	defer util.TimeTrack(time.Now(), "native-parallel")
	if len(a) == 0 || len(a[0]) == 0 || len(b) == 0 || len(b[0]) == 0 {
		return nil, fmt.Errorf("input matricies cannot be empty")
	}
	if len(a[0]) != len(b) {
		return nil, fmt.Errorf("incompattible matrix dimensions for multiplication")
	}

	result := make([][]float64, len(a))
	for i := range result {
		result[i] = make([]float64, len(b[0]))
	}

	var wg sync.WaitGroup

	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go func(rowIdx int) {
			defer wg.Done()
			for j := 0; j < len(b[0]); j++ {
				sum := 0.0
				for k := 0; k < len(b); k++ {
					sum += a[rowIdx][k] * b[k][j]
				}
				result[rowIdx][j] = sum
			}
		}(i)
	}

	wg.Wait()

	return result, nil
}

func MultiplyParallelWorkerPool(a, b [][]float64) ([][]float64, error) {
	defer util.TimeTrack(time.Now(), "native-parallel-worker-pool")
	if len(a) == 0 || len(a[0]) == 0 || len(b) == 0 || len(b[0]) == 0 {
		return nil, fmt.Errorf("input matrices cannot be empty")
	}
	if len(a[0]) != len(b) {
		return nil, fmt.Errorf("incompatible matrix dimensions")
	}

	result := make([][]float64, len(a))
	for i := range result {
		result[i] = make([]float64, len(b[0]))
	}

	numWorkers := runtime.NumCPU()

	jobs := make(chan int, len(a))

	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for rowIdx := range jobs {
				for j := 0; j < len(b[0]); j++ {
					sum := 0.0
					for k := 0; k < len(b); k++ {
						sum += a[rowIdx][k] * b[k][j]
					}
					result[rowIdx][j] = sum
				}
			}
		}()
	}

	for i := 0; i < len(a); i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()

	return result, nil
}
