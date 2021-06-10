package main

import "fmt"

func main() {

	/*  Like other programing languages, Go also supports all the major types of operators,
	Arithmetic Operators, Relational Operators, Logical Operators */

	x, y := 5, 8

	// Arithmetic
	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("y % x = ", y%x)

	// Relational Operators ->  "<, >, =, <=, >=" are relational operators.

	// Logical Operators
	code := true
	night := false

	fmt.Println("Code at night", code || night)
	fmt.Println("Code at daytime", code && night)

}
