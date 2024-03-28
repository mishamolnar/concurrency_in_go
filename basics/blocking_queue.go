package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type blockingQueue struct {
	cond   *sync.Cond
	arr    []int
	closed atomic.Bool
}

func newBlockingQueue() *blockingQueue {
	bq := blockingQueue{}
	bq.cond = sync.NewCond(&sync.Mutex{})
	return &bq
}

func (bq *blockingQueue) add(a int) {
	bq.cond.L.Lock()
	defer bq.cond.L.Unlock()
	bq.arr = append(bq.arr, a)
	bq.cond.Signal()
}

func (bq *blockingQueue) get() int {
	bq.cond.L.Lock()
	defer bq.cond.L.Unlock()
	for len(bq.arr) == 0 {
		bq.cond.Wait()
	}
	res := bq.arr[0]
	bq.arr = bq.arr[1:]
	return res
}

func try() {
	bq := newBlockingQueue()
	go func() {
		time.Sleep(5 * time.Second)
		bq.add(1)
	}()
	fmt.Println(bq.get()) //waits 5 second to get
}
