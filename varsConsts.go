package main

import "fmt"

func main() {

	// In Go declaration & initiallisation of variables and constants are of different types

	// keyword (var) <space> variable_name <space> DataType  - Basic Syntax
	const pi float32 = 3.14159

	var a int = 43
	var name string = "Abhinav-26"
	var abhi bool = true

	// For declaring multiple variables, Go is smart enough to understand the type of datatype provided in this type of declaration
	var (
		p = 10
		q = 11.4
	)

	// := is one of most widely used syntax for variables in go. It automatically understands the type of DataType used in this case
	x, y := 12, "abhinav"

	fmt.Println(pi)
	fmt.Println(a)
	fmt.Println(name)
	fmt.Println(abhi)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(x)
	fmt.Println(y)

}
