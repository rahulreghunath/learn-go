package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")
	var integ int = 10
	var ptr2 *int = &integ
	fmt.Println("Value of pointer2 is ", ptr2)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2

	fmt.Println(myNumber)
}
