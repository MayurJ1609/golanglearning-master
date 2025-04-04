package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// No inheritance in golang; No super or parent
	mayur := User{"Mayur", "mayur@go.dev", true, 31}
	fmt.Println(mayur)
	fmt.Printf("Mayur details are: %+v\n ", mayur)
	fmt.Printf("Name is %v and email is %v\n", mayur.Name, mayur.Email)
	mayur.GetStatus()
	mayur.NewMail()
	fmt.Printf("Name is %v and email is %v\n", mayur.Name, mayur.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (user User) GetStatus() {
	fmt.Println("Is user active: ", user.Status)
}

func (user User) NewMail() {
	user.Email = "test@go.dev"
	fmt.Println("email of this user is: ", user.Email)
}
