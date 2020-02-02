package main

import "fmt"

import "encoding/json"

// type person struct {
// 	First string `json:"firstname"`
// 	Last  string `json:"lastname,omitempty"`
// 	Age   int `json:",string"`
// }

type people []struct {
	Firstname string
	Lastname string
	Age int `json:",string"`
}

func main() {
	// p1 := person{
	// 	First: "Mohit",
	// 	Last:  "Sharma",
	// 	Age:   21,
	// }

	// p2 := person{
	// 	First: "Ron",
	// 	Last:  "Sharma",
	// 	Age:   21,
	// }

	// people := []person{p1, p2}
	// fmt.Println(people)
	// bs, err := json.Marshal(people)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(bs))

	jobject := `[{"Firstname":"Mohit","Lastname":"Sharma","Age":"21"},{"Firstname":"Ron","Lastname":"Bermer","Age":"21"}]`
	bs := []byte(jobject)
	
	var npeople people
	err := json.Unmarshal(bs, &npeople)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v",npeople[0])
}
