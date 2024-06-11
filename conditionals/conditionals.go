package conditionals

import (
	"errors"
)

func errorOut(msg string) (string, error) {
	if msg == "hello" {
		return "ok", nil
	}
	return "nope", errors.New("some error")
}

func LearnConditionals() {
	// if else
	// x := 100
	// if x > 10 {
	// 	fmt.Println("Greater than ten")
	// } else {
	// 	fmt.Println("nope <= 10")
	// }

	// shorthand -> first init then condition check
	if _, err := errorOut("hello"); err != nil {
		panic(err)
	}

	// if true {

	// } else if true {

	// } else {

	// }

	// a := "SPORTS"
	// b := "ANIME"
	// c := "SHERLOCK"

	// t := "spam" // network call -> received in response

	// // avoid nested if else
	// switch t {
	// case a:
	// 	fmt.Println("football")
	// case b:
	// 	fmt.Println("shippuden")
	// case c:
	// 	fmt.Println("Watson!")
	// default:
	// 	panic("invalid data")
	// }

	// // switch default true -> every case must evaluate to true.
	// switch {
	// case x < 10:
	// 	fmt.Println("football")
	// case x == 100:
	// 	fmt.Println("shippuden")
	// case x > 1000:
	// 	fmt.Println("Watson!")
	// default:
	// 	panic("invalid data")
	// }

}
