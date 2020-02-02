package exercises

import "fmt"

import "time"

import "sync"

import "math/rand"

func Mychann() {
	// normal channel and bi-directional channel
	// ch := make(chan int)
	// go func() {
	// 	ch <- 42
	// }()
	// fmt.Println(<-ch)

	// Buffer channel
	bch := make(chan int, 1)
	bch <- 23
	fmt.Printf("%T\n", bch)
	fmt.Println(<-bch)

	// Send only channel
	sc := make(chan<- int)
	fmt.Printf("%T\n", sc)

	// Receiver only channel
	rc := make(<-chan int)
	fmt.Printf("%T\n", rc)

	// bi-directional channel
	bdc := make(chan int)
	// not allowed
	// bdc = sc
	sc = bdc
	rc = bdc
	fmt.Printf("%T\n", sc)
	fmt.Printf("%T\n", rc)
}

func Testrs() {
	c := make(chan int)
	go sendDataToChannel(c)
	recvDataFromChannel(c)
	fmt.Println("I am here")
}

func sendDataToChannel(c chan<- int) {
	c <- 21
}

func recvDataFromChannel(c <-chan int) {
	fmt.Println(<-c)
}

func RangeChan() {
	c := make(chan int)
	go func() {
		for i := range c {
			fmt.Println(i)
		}
	}()
	for i := 0; i < 100; i++ {
		c <- i
	}
}

func RangeChanNew() {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
}

func TestNewOne() {
	// ec := make(chan int)
	// oc := make(chan int)
	// qc := make(chan bool)

	// go send(ec, oc, qc)
	// recv(ec, oc, qc)
	// c := fanin(boaring("Bob"), boaring("Mike"))
	// for i := 0; i < 30; i++ {
	// 	fmt.Println(i, <-c)
	// }

	// Fanout
	s := make([]int, 10)
	var wg sync.WaitGroup
	const noOfGoRoutines = 10
	for i := 1; i <= cap(s); i++ {
		s[i-1] = i
	}
	ic := make(chan int)
	for v := range s {
		go func(v int) {
			ic <- v
		}(v)
	}
	for i := 0; i < noOfGoRoutines; i++ {
		wg.Add(1)
		go func() {
			square(ic)
			wg.Done()
		}()
	}
	wg.Wait()
}

func square(c <-chan int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	v := <-c
	fmt.Println(v * v )
}

func send(e, o chan<- int, q chan bool) {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
	// q <- false
}

func recv(e, o <-chan int, q chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from even channel:", v)
		case v, ok := <-o:
			fmt.Println("from odd channel:", v, ok)
		case v, ok := <-q:
			if !ok {
				fmt.Println("from quit channel", v)
				fmt.Println("value of ok : ", ok)
				return
			}
		}
	}
}

func boaring(s string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- fmt.Sprint(s)
		}
	}()
	return c
}

func fanin(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()

	go func() {
		for {
			c <- <-c2
		}
	}()
	return c
}
