package exercises

import (
	"fmt"
	"os"
	"io/ioutil"
)

// TestError get exceptions.
func TestError() {
	var myname string
	_, err := fmt.Scan(&myname)
	if err != nil {
		fmt.Println(err)
	}
	_, err = fmt.Println("HI", myname)
	if err != nil {
		fmt.Println(err)
	}

	f, e := os.Create("test.txt")
	s := "Hello mohit"
	_, err = f.Write([]byte(s))
	if err != nil {
		fmt.Println(err)
	}

	f, e = os.Open("test.txt")
	if e != nil {
		fmt.Println(e)
	}
	bs , err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
