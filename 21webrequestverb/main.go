package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web verb")
	// performGetRequest()
	// performPostRequest()
	performPostFormRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code: ", response.StatusCode)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("details: ", string(content))

}

func performPostRequest() {
	const myurl = "http://localhost:8000/post"
	requestBody := strings.NewReader(`
	{
		"coursename":"MCA",
		"price":0,
		"platform":"MGU"
	}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func performPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "rahul")
	data.Add("lastName", "R")
	data.Add("address", "Sreesailam")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
