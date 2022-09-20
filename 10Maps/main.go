package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("The list of languages I know is", languages)
	fmt.Println("The value of JS is", languages["JS"])
	fmt.Println("I will delete some language")
	delete(languages,"PY")
	fmt.Println("New map is ", languages)
}