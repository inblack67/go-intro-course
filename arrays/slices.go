package arrays

import "fmt"

func LearnSlices() {
	// slices -> arrays with dynamic size -> we do not know the size of an array beforehand

	// intro
	// var arr []int
	// fmt.Println(arr == nil)

	// arr2 := []int{1, 2, 4}

	// [:]
	// making sub arrays
	// var arr = []int{1, 2, 3}
	// arr2 := arr[0 : len(arr)-1] // -> l : r -> r is not inclusive index
	// fmt.Println(arr2)

	// len(arr)
	// cap(arr)
	// Capacity: The number of elements in the underlying array starting from the index from which the slice is created.

	// append

	// arr = append(arr, 1344)

	// copy -> use it when new copy needed for slice. reassign will change the original slice unlike what happens in arrays
	source := []int{1, 2, 3}
	dest := source
	// dest := make([]int, len(source))
	// numberOfElementsCopied := copy(dest, source)
	dest[0] = 78
	fmt.Println(source)
	// fmt.Println(numberOfElementsCopied)
	fmt.Println(dest)

	// multi-dimensional
	// matrix := [][]int{}
}
