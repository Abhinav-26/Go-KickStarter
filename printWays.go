package main

import "fmt"

func main() {

	/* As like other programming languages, we can easily tweak what we want to print, and also it gives very simple methods
	for different use-cases */

	name := "Abhinav Dubey"
	x := 10
	pi := 3.15167890
	yes := true

	fmt.Println(len(name)) // it will give the length of string

	// Concatination of string using +
	fmt.Println(name + " is a chill dude.")

	// %f for floating numbers, can also assign the count of number we want after decimal like below code.
	fmt.Printf("%f \n", pi)
	fmt.Printf("%.3f \n", pi)

	// %d is used for integers
	fmt.Printf("%d \n", x)

	// %T is used to print the type of data type used.
	fmt.Printf("%T \n", name)

	// %t is used for boolean expressions
	fmt.Printf("%t \n", yes)

	// %b is used for binary number
	fmt.Printf("%b \n", 5)

	// %c is used to give character for the given number
	fmt.Printf("%c \n", 64)

	// %x is used to give hexcode of number
	fmt.Printf("%x \n", 15)

	// %e is used to print in scientific notation
	fmt.Printf("%e \n", pi)
}
