# Notes

### What is concurrency?
Concurrency is an ability to execute 2 or more proccess in the same time 

### What is race condition?
race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained

### Condition of deadlock?
Mutual Exclusion, Wait For Condition, No Preemption, Circular Wait

### What is livelock?
Livelocks are programs that are actively performing concurrent operations, but these operations do nothing to move the state of the program forward.

### What is starvation?
When more eager process takes resorce and leaves another process without it

### Difference between concurrency and parallelism? 
Concurrency is a property of the code; parallelism is a property of the running program.

### What is CSP?
“Communicating Sequential Processes.

### What will print this code?
```go
package main

import (
	"sync"
	"fmt"
)

func main() {
	str := "Bob"
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		str = "Alice"
	}()
	wg.Wait()
	fmt.Println(str) //what will print this?
}
```

it will print `Alice` because goroutine has an access to the outer scope

### How much space does goroutine consumes?
2-4 kb empty one:
we can calculate like this:
```go
func main() {
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
```

### Why to we need to call wait function on sync.WaitGroup in parent goroutine instead of calling it in child goroutine?
Because that way we create race condition and `hope` that this goroutine will be able to start. 
If it doesn't then waitGroup will not be blocked at all 


### Explain sync.Cond?
sync.Cond wraps the lock inside of it. It has 3 public methods to work with locks:
- `Wait()` - when calling wait lock inside Cond unlocks and thread becomes suspended. 
Should be called only withing lock. When `Wait` unblocks (returned) then lock becomes locked again 
- `Signal()` - wakes up goroutine that is waiting. 
- `Broadcast()` - wakes up all waiting goroutines
We can use sync.Cond to create blocking queue which waits till another thread adds to the queue 
