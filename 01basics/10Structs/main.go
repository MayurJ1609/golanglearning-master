package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// No inheritance in golang; No super or parent
	mayur := User{"Mayur", "mayur@go.dev", true, 31}
	fmt.Println(mayur)
	fmt.Printf("Mayur details are: %+v\n ", mayur)
	fmt.Printf("Name is %v and email is %v\n", mayur.Name, mayur.Email)

	s := User{"M1", "mayur@skorlife.com", true, 32}
	fmt.Println(s)
	fmt.Printf("S details are: %+v\n ", s)
	fmt.Printf("Name is %v and email is %v\n", s.Name, s.Email)

	sp := &s
	sp.Age = 28

	fmt.Println(sp)
	fmt.Println(s)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
