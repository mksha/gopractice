package main

import "fmt"

const (
	_ = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main1(){
	fmt.Printf("%b\t\n",kb)
	fmt.Printf("%b\t\n",mb)
	fmt.Printf("%b\t\n",gb)
	// a:= "H
	// fmt.Printf("%T\n",a)
	// fmt.Printf("%T\n",[]byte(a)[0])
	// fmt.Println([]byte(a)[0])
	// fmt.Printf("%b\n",[]byte(a)[0])
}

func newexercise(){
	x:= 42
	y:= "James Bond"
	z:= true
	fmt.Println(x,y,z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

