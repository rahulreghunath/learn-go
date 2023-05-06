package main

import "fmt"

func main() {
	fmt.Println("Struct")

	rahul := User{"Rahul", "me@rahulr.me", true, 27}
	fmt.Println(rahul)
	fmt.Printf("Details : %+v\n", rahul)
	fmt.Printf("Name: %v and email %v\n", rahul.Name, rahul.Email)
	rahul.GetStatus("message")
	rahul.NewMail()
	fmt.Printf("Name: %v and email %v\n", rahul.Name, rahul.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(message string) {
	fmt.Println("Is user active: ", u.Status)
	fmt.Println(message)
}

func (u User) NewMail() {
	u.Email = "rahul@gmail.com"
	fmt.Println("New email is ", u.Email)
}
