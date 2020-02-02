package exercises

import (
	"fmt"
	"sync/atomic"
	"runtime"
	"sync"
)
var wg sync.WaitGroup

func Level9(){
	
	fmt.Println("GOOS: ",runtime.GOOS)
	fmt.Println("GOARCH: ",runtime.GOARCH)
	fmt.Println("CPUs: ",runtime.NumCPU())
	fmt.Println("GoRutines:",runtime.NumGoroutine())

	wg.Add(1) // increment the the waitgroup counter
	go bar() // launching go routine to run bar
	foo()
	fmt.Println("GoRutines:",runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("GoRutines:",runtime.NumGoroutine())

}

func foo(){
	for i := 0 ; i < 10 ; i++ {
		fmt.Println("i am from foo : ",i)
	}
}

func bar(){
	defer wg.Done() // decrese the wg counter once go routine is completed
	for i := 0 ; i < 10 ; i++ {
		fmt.Println("i am from bar :",i)
	}
}

type person struct {
	Name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("I am from speak")
	fmt.Println(p.Name)
}


func saySomething(h human){
	h.speak()
}

func Level9_2(){
	p1 := person{Name: "Mohit"}
	saySomething(&p1)
}

func Level9_3(){
	var counter int64
	wg.Add(30)
	for i := 0; i < 30 ; i++{
		go func (){
			atomic.AddInt64(&counter, 1)
			defer wg.Done()
		}()
		fmt.Println("GoRoutines=", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(counter)

}
