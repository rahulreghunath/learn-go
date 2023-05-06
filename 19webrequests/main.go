package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web request")

	var url = "https://rahulr.me"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response  type is %T\n", response)
	// fmt.Println("Response: ",reresponse.)

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response :", string(databyte))
}
