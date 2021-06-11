package main

import "fmt"

func main() {

	// The concept of conditional statements, i.e - If else and Switch case are similar to other languages only.
	age := 10

	if age > 18 {
		fmt.Println("Congratulations! You can vote now")
	} else {
		fmt.Println("Oops! You cannot vote \n\n")
	}

	// Switch Case Statements
	switch age {
	case 10:
		fmt.Println("Hey kiddo! Study well")
	case 20:
		fmt.Println("Hey Dude! Great Going, you can vote")
	case 50:
		fmt.Println("Namaste Uncle! Please give your valuable vote")
	default:
		fmt.Println("Stay Healthy, Stay Safe")
	}
}
