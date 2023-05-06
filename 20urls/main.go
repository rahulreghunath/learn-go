package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?courcename=react"

func main() {
	fmt.Println("Urls")
	fmt.Println("My url: ", myUrl)

	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme, result.Host, result.Path, result.Port(), result.RawQuery)
	queryParams := result.Query()
	fmt.Printf("Type of queryparam %T\n", queryParams)

	fmt.Println(queryParams["courcename"])

	for _, value := range queryParams {
		fmt.Println("Value :", value)
	}

	partsOfUrl := url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=user",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
