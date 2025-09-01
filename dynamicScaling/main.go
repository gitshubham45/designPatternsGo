package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	tasks       chan int
	ctx         context.Context
	cancel      context.CancelFunc
	wg          sync.WaitGroup
	mu          sync.Mutex
	numWorkers  int
	minWorkers  int
	maxWorkers  int
	scaleTicker *time.Ticker
}

func NewWorkerPool(min, max int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		tasks:       make(chan int, 20),
		ctx:         ctx,
		cancel:      cancel,
		minWorkers:  min,
		maxWorkers:  max,
		scaleTicker: time.NewTicker(1 * time.Second), // monitor every second
	}
}

// worker logic
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for {
		select {
		case <-wp.ctx.Done():
			fmt.Printf("Worker %d shutting down\n", id)
			return
		case task, ok := <-wp.tasks:
			if !ok {
				fmt.Printf("Worker %d finished all tasks\n", id)
				return
			}
			fmt.Printf("Worker %d processing task %d \n", id, task)
			time.Sleep(500 * time.Millisecond) // simultate work
		}
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.minWorkers; i++ {
		wp.addWorker()
	}

	// Background sclaer
	go wp.autoScaler()
}

func (wp *WorkerPool) addWorker() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if wp.numWorkers < wp.maxWorkers {
		wp.numWorkers++
		id := wp.numWorkers
		wp.wg.Add(1)
		go wp.worker(id)
		fmt.Printf("worker %d started (total: %d) \n", id, wp.numWorkers)
	}
}

func (wp *WorkerPool) removeWorker() {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	if wp.numWorkers > wp.maxWorkers {
		wp.numWorkers--
		fmt.Printf("Scaling down - Remaining workers - %d \n", wp.numWorkers)
	}
}

func (wp *WorkerPool) autoScaler() {
	for {
		select {
		case <-wp.ctx.Done():
			return
		case <-wp.scaleTicker.C:
			queueLen := len(wp.tasks)

			if queueLen > wp.numWorkers && wp.numWorkers < wp.maxWorkers {
				wp.addWorker()
			} else if queueLen == 0 && wp.numWorkers > wp.maxWorkers {
				wp.removeWorker()
			}
		}
	}
}

func (wp *WorkerPool) submit(task int) {
	wp.tasks <- task
}

func (wp *WorkerPool) ShutDown() {
	fmt.Println("Shutting doen initiated...")
	wp.cancel()
	close(wp.tasks)
	wp.scaleTicker.Stop()
	wp.wg.Wait()
	fmt.Println("All workers stopped cleanly")
}

func main() {
	wp := NewWorkerPool(2, 6)
	wp.Start()

	for i := 1; i <= 20; i++ {
		wp.submit(i)
	}

	time.Sleep(10 * time.Second)

	wp.ShutDown()
}
