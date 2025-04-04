package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to creating JSON")
	// encodeJson()
	decodeJson()
}

func encodeJson() {
	lcoCourses := []course{
		{"React JS Bootcamp", 299, "youtube.com", "abc123", []string{"web", "dev", "js"}},
		{"MERN Bootcamp", 199, "youtube.com", "abcd123", []string{"web", "dev", "js"}},
		{"NESTJS Bootcamp", 399, "youtube.com", "abcde123", nil},
	}

	// Package data as JSON
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonData := []byte(`
	{
		"coursename": "React JS Bootcamp",
		"Price": 299,
		"website": "youtube.com",
		"tags": [
				"web",
				"dev",
				"js"
		]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON is not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and the value is %v and type is %T \n", k, v, v)
	}

}
