package exercises

// import (
// 	"fmt"
// 	"sort"
// 	"golang.org/x/crypto/bcrypt"
// )

// type person struct {
// 	Name string
// 	Age  int
// }

// type People []person

// type ByAge People

// type ByName People

// func (a ByAge) Len() int           { return len(a) }
// func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// func (a ByName) Len() int           { return len(a) }
// func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

// func Level8() {
// 	fmt.Println("Hi from level1")
// 	si := []int{9, 10, 11, 1, 22, 34, 4, 5, 60}
// 	ss := []string{"Hi", "Mohit", "Kumar", "Sharma", "How", "Are", "You"}
// 	fmt.Println(si)
// 	fmt.Println(ss)

// 	sort.Ints(si)
// 	sort.Strings(ss)

// 	fmt.Println(si)
// 	fmt.Println(ss)
// 	p1 := person{
// 		Name: "Mohit",
// 		Age:  23,
// 	}

// 	p2 := person{
// 		Name: "Narun",
// 		Age:  21,
// 	}

// 	peo1 := People{p1, p2}
// 	fmt.Println(peo1)

// 	sort.Sort(ByAge(peo1))
// 	fmt.Println(peo1)
// 	sort.Sort(ByName(peo1))
// 	fmt.Println(peo1)

// 	pass := `MohitKumar123@joy`
// 	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(passHash))
// 	err = bcrypt.CompareHashAndPassword(passHash, []byte(pass))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Password matched")
// }
