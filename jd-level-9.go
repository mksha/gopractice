package main

import "runtime"

import "fmt"

func main(){
	Level9_4()
}

func Level9_4(){
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}
