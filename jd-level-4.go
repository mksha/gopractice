package main

import "fmt"

func JdLevel4() {
	marray := [5]int{1, 2, 3, 4, 5}

	for i, v := range marray {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", marray)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", slice)

	nslice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(nslice[:5])
	fmt.Println(nslice[5:])
	fmt.Println(nslice[2:7])
	fmt.Println(nslice[1:6])

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:len(x)]...)
	fmt.Println(y)
	fmt.Println(x)
}

func jdl() {
	usn := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
	}
	fmt.Println(usn, cap(usn), len(usn))

	var nsum []string
	fmt.Println(nsum, cap(nsum), len(nsum))
	nsum = append(nsum, "a", "b", "c", "d")
	for i := 0; i < len(nsum); i++ {
		fmt.Println(i, nsum[i])
	}

	sl1 := []string{"James", "Bond", "Shaken"}
	sl2 := []string{"a", "b", "c"}
	mds := [][]string{sl1, sl2}
	fmt.Println(mds)
	for row, row_v := range mds {
		fmt.Println(row, row_v)
		for _, col_v := range row_v {
			fmt.Println(col_v)
		}
	}

	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "marantis", "women"},
		"moneypenny_miss": []string{"James Bond", "dance", "tea"},
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
	m["new_key"] = []string{"new", "time"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	delete(m, "bond_james")
	for k, v := range m {
		fmt.Println("new-----", k, v)
	}
}

func stru() {
	type people struct {
		Name    string
		Surname string
		Age     int
	}
	type secretAgent struct {
		people
		Id string
	}

	sa := secretAgent{
		people: people{
			Name:    "test1",
			Surname: "test2",
			Age:     32,
		},
		Id: "1234",
	}
	fmt.Println(sa)
	fmt.Println(sa.Name)

	people2 := struct {
		Name    string
		Surname string
		Age     int
	}{
		Name:    "test1",
		Surname: "test2",
		Age:     32,
	}
	fmt.Println(people2)
}

func jd5() {
	type person struct {
		first    string
		lastname string
		flavor   []string
	}

	p1 := person{
		first:    "test1",
		lastname: "test1-ln",
		flavor: []string{
			"vanila",
			"chocolate",
		},
	}
	p2 := person{
		first:    "test2",
		lastname: "test2-ln",
		flavor: []string{
			"mango",
			"apple",
		},
	}

	fmt.Println(p1, p2)
	for _, v := range p1.flavor {
		fmt.Println(v)
	}

	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
	}
}

func jd5new() {
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 1,
			color: "red",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(t1, s1)
}

func abc(a string) string {
	fmt.Println("Hi")
	return fmt.Sprint("Hi")
}

func xyz(a string , b int, c bool) (string, int, bool){
	return a,b,c
}

func sum(n ...int) int {
	sum := 0
	fmt.Println(n)
	fmt.Printf("%T\n",n)
	fmt.Println(cap(n))
	fmt.Println(len(n))

	for _ , v := range n {
		sum += v
	}
	return sum
}

type agent struct {
	fname string
	lname string
}

type secret struct {
	choice string
}
func (a agent) cprint() {
	fmt.Println("Hi ")
}
func (a secret) cprint() {
	fmt.Println("Hi ")
}
type human interface {
	cprint()
}

func nn(h human){
	switch h.(type){
	case agent:
		fmt.Println(h.(agent).fname)
	case secret:
		fmt.Println(h.(secret).choice)
	}
	fmt.Println("k",h)
}
func main4() {
	// s := []int{1,2,3,4,5}
	// fmt.Println(sum(s...))
	// defer abc("Mohit")
	// fmt.Println(xyz("jit",1,true))
	a1 := agent{
		fname : "Mohit",
		lname : "Sharma",
	}
	s1 := secret{
		choice: "my choice",
	}

	a1.cprint()
	nn(a1)
	nn(s1)
	fmt.Printf("%T\n",a1)

}
