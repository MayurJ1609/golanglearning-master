package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ksdjfksjfeiru"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("Type of query params are %T\n", queryParams)
	fmt.Println(queryParams["coursename"])

	for _, v := range queryParams {
		fmt.Println("Param is: ", v)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=mayur",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
