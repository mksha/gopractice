package main

import "fmt"

func main2() {
	// const placeOfInterest = "\xdd"
	// fmt.Printf("%s\n",placeOfInterest)
	// fmt.Printf("%+q\n",placeOfInterest)
	// for i:=0; i < len(placeOfInterest); i++{
	// 	fmt.Printf("index=%v\t%x\n",i, placeOfInterest[i])
	// }
	// s := "Hello mohit"
	// fmt.Printf("%T\n", s[0])
	// sb := []byte(s)
	// fmt.Println(sb)
	
	// for index, value := range(s) {
	// 	fmt.Println(index,value)
	// 	fmt.Printf("%U\n",value)
	// }
	// const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// sb := []byte(sample)
	// for i:=0; i < len(sb); i++{
	// 	fmt.Printf("%+q\n",sb[i])
	// }
	// fmt.Println(sample)
	// fmt.Printf("%+q\n",sample)
	// for i:=0 ; i < len(sample); i++ {
	// 	fmt.Printf("%x ", sample[i])
	// }
	// // exercise2()
	// fmt.Println(runtime.GOOS)
	// fmt.Println(runtime.GOARCH)
	// func() {
	// 	fmt.Println("Hi Mohit")
	// }()

	// func(x int) {
	// 	fmt.Println("Hi Mohit", x)
	// }(2)

	// funcexp := func(x int) {
	// 	fmt.Println("This is from func expression", x)
	// }
	// funcexp(21)

	// f1 := foo1()
	// fmt.Printf("%T\n",f1)
	// fmt.Println(f1())
	// slice := []int {1,2,3,4,5,6,7,8}
	// fmt.Println(evenOdd(sumn, slice...))

	// inc := incr()
	// dec := decr(5)
	// fmt.Println(inc())
	// fmt.Println(inc())
	// fmt.Println(dec())
	// fmt.Println(dec())
	fmt.Println(fact(5))
	
}

func fact(n int) int{
	if n == 0 {
		return 1
	}else {
		return n*fact(n-1)
	}
}
func incr() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func decr(v int) func() int {
	return func() int {
		v--
		return v
	}
}

func foo1() func() int {
	return func() int {
		return 21
	}
}

func sumn(slice ...int) int {
	sum := 0 
	for _ , v := range slice {
		sum += v
	}
	return sum
}

func evenOdd(f func(slice1 ...int) int ,slice ...int) (int , int){
	es := []int{}
	os := []int{}
	for _ , v := range slice {
		if v % 2 == 0 {
			es = append(es, v)
		}else {
			os = append(os,v)
		}
	}
	ESum := f(es...)
	OSum := f(os...)
	return ESum,OSum
}
