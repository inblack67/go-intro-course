package arrays

import "fmt"

func LearnArrays() {
	// array intro
	// arr := [5]int{101, 2, 3, 4, 5}
	// arrStr := [5]string{"abc"}
	// fmt.Println(arr[0])
	// fmt.Println(len(arr))

	// // init with size

	// // ...
	// ok := [...]int{1, 2, 34, 4} // golang will calculate the size itself
	// fmt.Println(ok)

	// loop over
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// for _, v := range arr {
	// 	fmt.Println(v)
	// }

	// copyinh -> array as value types
	var arr1 = [...]int{1, 2, 4}
	arr2 := arr1

	arr2[0] = 567

	// fmt.Println(arr1[0])
	// fmt.Println(arr2[0])

	// 2d array

	matrix := [5][5]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}

	fmt.Println(matrix[0][0])

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			fmt.Println(matrix[row][col])
		}
	}

}
