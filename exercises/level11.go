package exercises

import "fmt"
import "runtime"
import "sync"

func Level11() {
	var wg sync.WaitGroup
	fmt.Println("Goroutine=",runtime.NumGoroutine())
	wg.Add(1)
	go func() {
		fmt.Println("Goroutine=",runtime.NumGoroutine())
		fmt.Println("i am from diff function")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Goroutine=",runtime.NumGoroutine())
		fmt.Println("i am from diff function")
		wg.Done()
	}()
	wg.Wait()
}

