package main

import (
	"encoding/json"
	"fmt"
)

type cource struct {
	Name     string `json:"courcename"`
	Price    int    `json:"price"`
	Platform string `json:"-"`
	Password string
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	cources := []cource{
		{"react name", 299, "Learbcodeonline", "abc1231", []string{"web-dev", "js"}},
		{"react name2", 200, "Learbcodeonline", "abc1232", []string{"web-dev", "js"}},
		{"react name3", 500, "Learbcodeonline", "abc1233", nil},
	}

	// package this  data as json

	finalJson, err := json.MarshalIndent(cources, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"courcename": "react name",
			"price": 299,
			"Password": "abc1231",
			"tags": [
				"web-dev",
				"js"
			]
		}
	`)
	var lcoCource cource
	is_valid := json.Valid(jsonDataFromWeb)
	if is_valid {
		fmt.Println("Valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCource)
		fmt.Printf("%#v\n", lcoCource)
	} else {
		fmt.Println("JSON NOT VALID")
	}

	// to add data to add data to key

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("\n\n%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key: %v and value %v and type is %T \n", k, v, v)
	}

}
