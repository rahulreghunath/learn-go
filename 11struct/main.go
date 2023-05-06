package main

import "fmt"

func main() {
	fmt.Println("Struct")

	rahul := User{"Rahul", "me@rahulr.me", true, 27}
	fmt.Println(rahul)
	fmt.Printf("Details : %+v\n", rahul)
	fmt.Printf("Name: %v and email %v\n", rahul.Name, rahul.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
