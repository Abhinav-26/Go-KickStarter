package main

import "fmt"

func main() {

	// Like c/c++ we have Pointers in golang as well.

	x := 10

	fmt.Println(x)
	fmt.Println(&x) // & represents the address of varibale x

	// USE-CASE - if we want to change the value of certain variable directly form the memory location, pointer comes into play

	changeValue(&x)
	fmt.Println(x) // this will print the chnaged value, because value has been directly changed from the address
}

func changeValue(x *int) {
	*x = 7
}
