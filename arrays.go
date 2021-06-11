package main

import "fmt"

func main() {

	// the concept is same as of other languages, but the syntax is quite different
	arr := [5]int{1, 2, 4, 5, 6}
	fmt.Println(arr)

	// to get from specific index
	fmt.Println(arr[3])

	// make creates a new blank array of size as specified 5,
	new_arr := make([]int, 5)
	// Will return a blank array with size 5
	fmt.Println(new_arr)

	// for iterating over array
	for _, value := range arr {
		fmt.Println(value)
	}
}
