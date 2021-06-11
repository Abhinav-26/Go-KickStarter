package main

import "fmt"

func main() {

	main_arr := []int{1, 2, 3, 4, 5}

	// make creates a new blank array of size as specified 5,
	arr := make([]int, 5)
	fmt.Println(main_arr)
	fmt.Println(arr, "\n")

	// copy methods copied the second array into the first one passed in arguments.
	copy(arr, main_arr)
	fmt.Println(main_arr)
	fmt.Println(arr)

	// Slicing methods, in slicing like python, the last index element is never considered
	fmt.Println(arr[1:3])

	// arr[:0] or arr[0:] - in these cases it specifies everything after or before the index passed beside to :, same like python
	fmt.Println(arr[0:])
	fmt.Println(arr[:4])

	new_arr := append(arr, 9, 8, 7)
	fmt.Println(new_arr)
}
