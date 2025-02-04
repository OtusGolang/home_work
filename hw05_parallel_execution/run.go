package hw05parallelexecution

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// RunHw5 starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func RunHw5(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return errors.New("no tasks to run")
	}
	var countErr int32 = 0
	var countParallelGo = int32(n)
	taskChan := make(chan Task, len(tasks))
	for _, task := range tasks {
		taskChan <- task
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := sync.WaitGroup{}

	for {
		select {
		case <-ctx.Done():
			return ErrErrorsLimitExceeded
		case <-taskChan:
			if atomic.LoadInt32(&countParallelGo) <= 0 {
				fmt.Println("not finished")
				continue
			}
			if task, ok := <-taskChan; ok {
				//was sleep
				wg.Add(1)
				atomic.AddInt32(&countParallelGo, 1)
				val := atomic.LoadInt32(&countErr)
				go func() {
					defer wg.Done()
					fmt.Println("LOADED VAL", val)
					if val >= int32(m) {
						cancel()
					}
					err := task()
					if err != nil {
						fmt.Println("ERROR", err)
						atomic.AddInt32(&countErr, 1)
					}
					atomic.AddInt32(&countParallelGo, -1)
				}()
			} else {
				return nil
			}
		case <-time.After(25 * time.Second):
			return errors.New("timeout")
		}

		wg.Wait()
	}
}
