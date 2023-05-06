package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("googe")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
