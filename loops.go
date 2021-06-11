package main

import "fmt"

func main() {

	// like other programming languages, the concept of loops are similar in golang as well.
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// exit control loop (while loop approach)
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	/* The basic syntax and concept of loops is this, similary we can use the same concepts in nested loops and in other
	usecases as per requirements */
}
