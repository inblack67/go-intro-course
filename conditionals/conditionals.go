package conditionals

import "fmt"

func LearnConditionals() {
	// if else
	x := 100
	if x > 10 {
		fmt.Println("Greater than ten")
	} else {
		fmt.Println("nope <= 10")
	}

	// if true {

	// } else if true {

	// } else {

	// }

	a := "SPORTS"
	b := "ANIME"
	c := "SHERLOCK"

	t := "spam" // network call -> received in response

	// avoid nested if else
	switch t {
	case a:
		fmt.Println("football")
	case b:
		fmt.Println("shippuden")
	case c:
		fmt.Println("Watson!")
	default:
		panic("invalid data")
	}
}
