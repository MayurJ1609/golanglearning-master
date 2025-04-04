package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const BaseUrl string = "http://localhost:8000/"

func main() {
	fmt.Println("Welcome to web requests")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	myUrl := BaseUrl + "postform"

	// formdata
	data := url.Values{}
	data.Add("fn", "mayur")
	data.Add("sn", "j")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformPostJsonRequest() {
	myUrl := BaseUrl + "post"

	// fake JSON payload
	requestBody := strings.NewReader(`
		{
			"courseName": "Go with Golang",
			"price": 100
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func PerformGetRequest() {
	myUrl := BaseUrl + "get"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))

}
