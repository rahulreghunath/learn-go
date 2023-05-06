package main

import "fmt"

func main() {
	fmt.Println("If else")

	loginCount := 25
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		fmt.Println("watcth out")
	} else {
		result = "Something"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Greater than 10")
	}

}
