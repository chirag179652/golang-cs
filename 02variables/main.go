package main

import "fmt"

const LoginToken string = "z34i3n43da434b34a34d"

func main() {
	var username string = "chirag17"
	fmt.Println(username)
	fmt.Printf("variable is of type %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type %T \n", isLoggedIn)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type %T \n", LoginToken)

}