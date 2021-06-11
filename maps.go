package main

import "fmt"

func main() {

	// maps in golang is some what similar to python dictionary, it declaration syntax is -
	StudentAge := make(map[string]int)

	// KeyValue pair concept
	StudentAge["Abhinav"] = 20
	StudentAge["Abhi"] = 19
	StudentAge["Kaustubh"] = 22
	StudentAge["Prasann"] = 24
	StudentAge["Aman"] = 21

	fmt.Println(StudentAge)

	// to print a specific value
	fmt.Println(StudentAge["Abhinav"])
}
