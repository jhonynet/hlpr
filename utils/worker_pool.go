package utils

import (
	"sync"
)

type WorkerPool struct {
	workers chan int
	wg      *sync.WaitGroup
}

func NewWorkerPool(workers int) *WorkerPool {
	w := &WorkerPool{
		workers: make(chan int, workers),
		wg:      &sync.WaitGroup{},
	}
	return w.fill(workers)
}

func (w *WorkerPool) Work(function func()) {
	w.wg.Add(<-w.workers)
	go func() {
		defer w.wg.Done()
		defer func() { w.workers <- 1 }()
		function()
	}()
}

func (w *WorkerPool) Wait() *WorkerPool {
	w.wg.Wait()
	return w
}

func (w *WorkerPool) Close() *WorkerPool {
	close(w.workers)
	return w
}

func (w *WorkerPool) fill(workers int) *WorkerPool {
	for t := 0; t < workers; t++ {
		w.workers <- 1
	}
	return w
}
