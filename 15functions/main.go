package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greeter()

	result := adder(3, 5)
	fmt.Printf("Result is %v\n", result)

	result2, message := proAdder(3, 5, 3, 4, 4)
	fmt.Printf("Result is %v, message is %v\n", result2, message)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hi Pro result function"
}

func greeter() {
	fmt.Println("Namasthe from go")
}
