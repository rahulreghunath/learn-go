package main

import "fmt"

const LoginToken string = "dfsdfsdfs" // Public variable

func main() {
	var username string = "Rahul"
	fmt.Println(username)
	fmt.Printf("Variable  is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable  is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable  is of type: %T \n", smallVal)

	var smallFloat float64 = 255.23423423423242332323234
	fmt.Println(smallFloat)
	fmt.Printf("Variable  is of type: %T \n", smallFloat)

	// Default values and aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable  is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable  is of type: %T \n", anotherString)

	// implicit type
	var website = "rahulr.me"
	fmt.Printf("Variable  is of type: %T \n", website)
	fmt.Println(website)

	// no var style
	numberOfUser := 30000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable  is of type: %T \n", LoginToken)
}
