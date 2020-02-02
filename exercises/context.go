package exercises

import "context"

import "fmt"

import "runtime"

import "time"

func Context() {
	emptyContext := context.Background()

	fmt.Println(emptyContext)
	fmt.Printf("%T\n", emptyContext)
	fmt.Println(emptyContext.Err())

	copyOfParentContext, cancel := context.WithCancel(emptyContext)

	fmt.Println(copyOfParentContext)
	fmt.Printf("%T\n", copyOfParentContext)
	fmt.Println(copyOfParentContext.Err())
	fmt.Println("-----------------------------")
	// copyOfParentContext2, cancel := context.WithTimeout(emptyContext, time.Minute*5)
	// fmt.Println(copyOfParentContext2)
	// fmt.Printf("%T\n", copyOfParentContext2)
	// fmt.Println(copyOfParentContext2.Err())
	// cancel()
	// fmt.Println(copyOfParentContext2)
	// fmt.Printf("%T\n", copyOfParentContext2)
	// fmt.Println(copyOfParentContext2.Err())

	// fmt.Println("-----------------------------")
	// copyOfParentContext3, cancelFunc := context.WithDeadline(emptyContext, time.Unix(60, 60))
	// fmt.Println(copyOfParentContext3)
	// fmt.Printf("%T\n", copyOfParentContext3)
	// fmt.Println(copyOfParentContext3.Err())
	// cancelFunc()
	// fmt.Println(copyOfParentContext3)
	// fmt.Printf("%T\n", copyOfParentContext3)
	// fmt.Println(copyOfParentContext3.Err())
	// fmt.Println(runtime.NumGoroutine())
	go func() {
		i := 0
		fmt.Println("from go routine 2:", runtime.NumGoroutine())
		for {
			select {
			case <-copyOfParentContext.Done():
				return
			default:
				i++
				fmt.Println(i)
			}
		}
	}()
	fmt.Println("from main ", runtime.NumGoroutine())
	cancel()
	time.Sleep(time.Second * 2)
	fmt.Println("from main ", runtime.NumGoroutine())
	fmt.Println(copyOfParentContext.Err())
}

func Learning() {
	for v := range sq(gen(2, 3)) {
		fmt.Println(v)
	}
}

func gen(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range num {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}
