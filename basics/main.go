package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)


func calcMemoryPerGoroutine() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan interface{}
	var wg sync.WaitGroup
	fn := func() {
		wg.Done()
		<-c
	}
	const numRoutines = 1e4
	wg.Add(numRoutines)
	before := memConsumed()
	for i := numRoutines; i > 0; i-- {
		go fn()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb \n", float64(after-before)/numRoutines/1000)
	fmt.Printf("init memory %.3fkb \n", float64(before)/1000)
	fmt.Printf("after memory %.3fkb \n", float64(after)/1000)
}

type num struct {
	n  int
	mu sync.Mutex
}

func deadlock() {
	var a, b num
	var wg sync.WaitGroup
	printSum := func(a, b *num) {
		defer wg.Done()
		a.mu.Lock()
		i := a.n
		defer a.mu.Unlock()

		time.Sleep(2 * time.Second)

		b.mu.Lock()
		j := b.n
		defer b.mu.Unlock()

		fmt.Println(i + j)
	}
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
