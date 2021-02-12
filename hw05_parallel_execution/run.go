package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n int, m int) error {
	taskCh := make(chan Task, len(tasks))

	go func() {
		defer close(taskCh)

		for _, task := range tasks {
			taskCh <- task
		}
	}()

	var errCounter int32

	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for task := range taskCh {
				if err := task(); err != nil {
					atomic.AddInt32(&errCounter, 1)
				}

				if m != 0 && int(atomic.LoadInt32(&errCounter)) >= m {
					break
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	if m != 0 && int(errCounter) >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
