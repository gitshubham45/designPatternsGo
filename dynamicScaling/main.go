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

func NewWorkerPool(min, max int) *NewWorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &NewWorkerPool{
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
			fmt.Println("Worker %d processing task %d \n", id, task)
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
		go wp.
	}
}

func main() {
	wp := NewWorkerPool(2, 6)

}
