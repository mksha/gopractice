package exercises

import (
	"fmt"
	"log"
	"os"
)

func TestLogging() {
	f, err := os.Create("test.log")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)
	_, err = os.Open("no-file")
	if err != nil {
		// log.Println(err)
		// log.Fatalln(err)
		// log.Panicln(err)
		panic(err)
	}
}
