package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println("List all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["js"])

	delete(languages, "rb")
	fmt.Println("List all languages: ", languages)

	// Loops in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
