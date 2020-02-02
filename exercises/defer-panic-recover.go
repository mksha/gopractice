package exercises

import "fmt"

// import "errors"

// import "log"

import "errors"

// NegativeSquareError type
type NegativeSquareError struct {
	data float64
	err  error
}

type customErr struct {
	data string
	err error
}

func foosample(e error){
	fmt.Println(e.(customErr).data)
}

func (ce customErr) Error() string {
	return fmt.Sprintf("%v",ce.err)
}

func (nse NegativeSquareError) Error() string {
	return fmt.Sprintf("NegativeSquareError : Can not find the square root of Negative number %f \n  %v", nse.data, nse.err)
}

func DeferPanicRecover() {
	// fmt.Println("I am starting execution of DeferPanicRecover")
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Panic was captured", r)
	// 	}
	// }()
	// footest(0)
	// fmt.Println("All set")
	// v, err := sqroot(-2)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(v)
	err := errors.New("Killing")
	ce := customErr{"HI",err}
	foosample(ce)
}

func sqroot(n float64) (float64, error) {
	if n < 0 {
		// err := errors.New("Invalid Values : Can not find the square root of negative number")
		err := fmt.Errorf("Invalid Value : %f", n)
		return 0, NegativeSquareError{2, err}
	}
	return n * n, nil
}

func footest(i int) {
	if i > 3 {
		fmt.Println("value of i", i)
		panic(i)
	}
	defer fmt.Println("I am from defer", i)
	fmt.Println("Outside of if branch")
	footest(i + 1)
	fmt.Println("After foo call in foo")
}
