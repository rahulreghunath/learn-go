package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating:")

	// coma ok || error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Printf("Type of this rating, %T", input)
}
