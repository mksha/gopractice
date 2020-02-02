package main

import "fmt"

const (
	_ = iota
	a = 2019 + iota
	b = 2019 + iota
)

func jd2() {
	x := 15
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
}

func main3() {
	// jd2()
	// for i := 0; i < 3; i++ {
	// for j := 0; j < 3; j++ {
	// fmt.Printf("outer loop index=%d\t inner loop index=%+q\n",i,j)
	// fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	// fmt.Printf("\t%#U\n", i)
	// }
	// for i := 33; i <= 122; i++ {
	// 	fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	// }
	// }

	// for k:= 30; k< 122 ; k++ {
	// 	fmt.Printf("%#U\n",k)
	// }
	// switch x := "Hi"; x {
	// case "Nn":
	// 	fmt.Println("case 1")
	// case "Hi":
	// 	fmt.Println("case 2")
	// case "nnc":
	// 	fmt.Println("case 3")
	// default:
	// 	fmt.Println("default")
	// }

	// switch z := 1; c := x.(type) {
	// case nil:
	// 	fmt.Println("nil")
	// case int:
	// 	fmt.Println("int")
	// case string:
	// 	fmt.Println("string")
	// }
	// exercise()
	// fmt.Println(true && true)
	// fmt.Println(true && false)
	// fmt.Println(false && true)
	// fmt.Println(true || true)
	// fmt.Println(true || false)
	// fmt.Println(false || false)
	maptest()
	
}

func even(n int) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Println("Number is even", i)
		} else {
			fmt.Println("Number is not even", i)
		}
	}
}

func exercise() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
	year := 1996
	for {
		if year <= 2020 {
			fmt.Println(year)
			year++
		} else {
			break
		}
	}

	for n := 10 ; n <= 100 ; n++ {
		fmt.Println(n % 4)
	}

	switch favSpot := "cricket" ; favSpot {
	case "cricket":
		fmt.Println("cricket")
	default:
		fmt.Println("default")
	}

}

func testarray() {
	var myarray [4]int
	fmt.Println(myarray)
	fmt.Println(len(myarray))
	fmt.Printf("%T\n",myarray)

	x := []int{1,2,3,4,5,6,7,8}
	y := append(x[:2],x[5:]...)
	fmt.Println(y)
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	for i, v := range x {
		fmt.Println(i,v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(x[0:i])
	}
	var uis []int
	fmt.Println(uis)
}

func muname(){
	var array [4]int
	var slice2 = array[:len(array)-1]
	fmt.Println(slice2)
	myslice := make([]int , 10, 20)
	fmt.Println(myslice)
	fmt.Println(append(slice2,myslice...))
	fmt.Println(append(slice2[:1],slice2[2:]...))
	myslice = myslice[:len(myslice)+1]
	slice2 = slice2[:len(slice2)+1]
	fmt.Println(myslice)
	myslice[10] = 54
	fmt.Println(myslice)
}

func slicemulti(){
	slice := []string{"test","new","feature"}
	mslice := [][]string{slice,slice}
	type ms [][]string
	bslice := ms{slice,slice}
	fmt.Println(slice)
	fmt.Println(mslice)
	fmt.Println(bslice)

}

func maptest(){
	map1 := map[string]int{
		"Name": 1,
		"Purpose": 2,
	}
	fmt.Println(map1)
	fmt.Println(map1["Name"])
	if v, ok := map1["hit"] ; ok {
		fmt.Println(v)
	}
	nmap := make(map[string]int,2)
	nmap["name"]=12
	nmap["name2"]=23
	nmap["name3"]=45
	fmt.Println(nmap)
	for k, v := range nmap {
		fmt.Println(k,v)
	}
	delete(nmap,"name")
	fmt.Println(nmap)
	testn := [...]int{1,2,3,4,4}
	fmt.Printf("%T\n",testn)
}
