package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Orange"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"Potato", "beans", "mushroom"}
	fmt.Println("Vegi list: ", vegList)
	fmt.Println("Vegi list length: ", len(vegList))
}
