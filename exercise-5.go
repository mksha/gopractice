package main

import "fmt"

// func createPointer(){
// 	a := 2
// 	fmt.Println(a)
// 	fmt.Printf("%T\n",a)
// 	fmt.Println(&a)
// 	fmt.Printf("%T\n",&a)

// 	b := &a
// 	fmt.Println(b)
// 	fmt.Println(*b)
// 	fmt.Printf("%T\n",b)

// 	*b = 21
// 	fmt.Println(a)

// }

func tap(ptr *int) {
	*ptr = 21
}

type person struct {
	Name string
	Age int
}
func (d person) changeMe(p *person){
	fmt.Println(p)
	// *(&p.Age) = 21
	// (*p).Age = 21
	p.Age = 21
}
func main() {
	// createPointer()
	// x := 2
	// tap(&x)
	// fmt.Println(x)

	p1 := person{
		Name: "Mohit",
		Age: 20,
	}

	fmt.Println(&(p1.Age))
	fmt.Println(&p1.Age)

	// changeMe(p1)
	(&p1).changeMe(&p1)
	fmt.Println(p1)
}
