package datatypes

import (
	"fmt"
)

func LearnDatatypes() {
	// int -> 32, int32, int64, rune -> int32

	var x int // 0 by default, vars whenever we wanna explicitly give type to variable or declare a global variable
	x = 10
	fmt.Println(x)

	// shorthand -> widely used
	y := 10 // type inference
	fmt.Println(y)

	// float -> float32, 64 ...
	f := 10.54
	fmt.Println(f)

	// strings
	str := "aman"
	fmt.Println(str) // loop over -> in loops

	// composite data types arrays, slices, structs -> in later videos
}
