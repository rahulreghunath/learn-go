package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case")
	rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
	case 3:
		fmt.Println("Value is 3")
	case 4:
		fmt.Println("Value is 4")
	case 5:
		fmt.Println("Value is 5")
	case 6:
		fmt.Println("Value is 6")
		fmt.Println("One more chance")

	default:
		fmt.Println("Wjat was that")
	}
}
