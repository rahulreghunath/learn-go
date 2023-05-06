package main

import "fmt"

func main() {
	fmt.Println("Loops")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, value := range days {
	// 	fmt.Printf("index is and value %v\n", value)
	// }
	rougueValue := 1

	for rougueValue < 10 {
		fmt.Println("Value is", rougueValue)
		rougueValue++

	}
}
