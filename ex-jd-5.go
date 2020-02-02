package main

import "fmt"

func foo(slice ...int) int {
	sum := 0
	for _ , v := range slice {
		sum += v
	}
	return sum
}

func bar(slice []int) int{
	sum := 0
	for _ , v := range slice {
		sum += v
	}
	return sum
}

func ntest(){
	fmt.Println("Hi this is defer")
}

type person struct {
	first string
	last string
	age int
}

func (p person) speak(){
	fmt.Println(p.first, p.age)
}

type SQUARE struct {
	L float64
	W float64
}

type CIRCLE struct {
	R float64
}

func (c CIRCLE) area() float64 {
	r := c.R
	return 3.14 * r * r
}

func (s SQUARE) area() float64 {
	return s.L * s.W
}

type SHAPE interface {
	area() float64
}

func info(s SHAPE){
	// switch s.(type){
	// case SQUARE:
	// 	fmt.Println(s.(SQUARE).area())
	// case CIRCLE:
	// 	fmt.Println(s.(CIRCLE).area())
	// }
	fmt.Println(s.area())

	f := func(){
		fmt.Println("I am anonymous function")
	}
	f()
}

func main(){
	// slice := []int{1,2,3,4,5}
	// defer ntest()
	// fmt.Println(foo(slice...))
	// fmt.Println(bar(slice))
	s1 := SQUARE{
		L: 12.0,
		W: 10.5,
	}
	c1 := CIRCLE{
		R: 10.0,
	}
	fmt.Println("ARea of Square=",s1.area())
	info(s1)
	fmt.Println("ARea of Circle=",c1.area())
	info(c1)

	// p1 := person{
	// 	first: "Mohit",
	// 	last: "Sharma",
	// 	age: 21,
	// }
	// p1.speak()
	fa := rf()
	fmt.Println(fa())
}

func rf() func() string {
	return func() string {
		return "Hi Ano"
	}
}
