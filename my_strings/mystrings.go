package mystrings

import (
	"fmt"
	"strconv"
)

func LearnStrings() {
	// global = "hello worlds"
	// // strings
	// // var str string = "I love Elixir!"
	// // fmt.Println(global)
	// str := "a go world"

	// for i := 0; i < len(str); i++ {
	// 	// fmt.Println(str[i]) // ASCII values of chars
	// 	// fmt.Printf("%c \n", str[i]) // chars
	// 	c := fmt.Sprintf("%c", str[i]) // store char in var
	// 	fmt.Println(c)
	// }

	// // split string
	// strArr := strings.Split(str, " ")
	// fmt.Println(strArr)
	// fmt.Println(strArr[0])

	// fmt.Println(strings.ToUpper(strArr[0]))

	// array of strings
	// arrStr := []string{"ts", "go", "ex", "rs"}

	// looping through
	// for i := 0; i < len(arrStr); i++ {
	// 	// fmt.Println(arrStr[i])
	// }

	// shorthand loop -> range
	// for i, v := range arrStr {
	// 	fmt.Printf("index=%d, value=%s\n", i, v)
	// }

	// arrStr2 := make([]string, 5) // defaults to empty string for every element uptil size declared
	// fmt.Println(arrStr2[0] == "")
	// fmt.Println(arrStr[0])

	// fmt.Println(len("lengthy string")) // length of chars in string

	// // slices
	// res := []string{}

	// // fmt.Println((res)) // length of elements in array/slice of strings

	// // push
	// res = append(res, "I am first")
	// res = append(res, "I am mid")
	// res = append(res, "I am last")

	// // pop
	// res = res[1 : len(res)-1]

	// fmt.Println(res)

	// conversions -> strconv
	numStr := "12c" // need it as an integer
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic("invalid string of int")
	}
	// strconv.Itoa()
	fmt.Println(num)
}
