package exercises

import "fmt"
import "sync"
import "sync/atomic"
import "runtime"


func Level10(){
	var counter int64
	gs := 100
	var wg sync.WaitGroup
	// var me sync.Mutex
	wg.Add(gs)

	for i :=0 ; i < gs ; i++ {
		go func () {
			// me.Lock()
			atomic.AddInt64(&counter, 1)
			// counter++
			atomic.LoadInt64(&counter)
			// me.Unlock()
			// fmt.Println("Goroutine:",runtime.NumGoroutine())
			// runtime.Gosched()
			// counter = v
			wg.Done()
		} ()
		fmt.Println("Goroutine:",runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
