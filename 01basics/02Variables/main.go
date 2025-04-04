package main

import "fmt"

const LoginToken string = "kajslkuri3943ijkdkfk" // Public variable

func main() {
	fmt.Println("Variables")

	// Type1: using var keyword and initializing it (mentioning type)
	var username string = "Mayur"
	fmt.Println(username)
	fmt.Printf("Variable of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable of type: %T \n", smallVal)

	var smallFloat float32 = 255.2372637267642736273672
	fmt.Println(smallFloat)
	fmt.Printf("Variable of type: %T \n", smallFloat)

	fmt.Println(("--------------------------------------"))
	// Type2: Default values ans some aliases (mentioning type)
	var anotherIntVariable int
	fmt.Println(anotherIntVariable)
	fmt.Printf("Variable of type: %T \n", anotherIntVariable)

	var anotherStringVariable string
	fmt.Println(anotherStringVariable)
	fmt.Printf("Variable of type: %T \n", anotherStringVariable)
	fmt.Println(("--------------------------------------"))

	// Type3: implicit type (Not mentioning type)

	var website = "mayur.com"
	fmt.Println(website)
	fmt.Printf("Variable of type: %T \n", website)

	// Type4: No Var style (Not mentioning type). Cannot be used outside function

	numberOfUsers := 3000000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable of type: %T \n", LoginToken)

}
