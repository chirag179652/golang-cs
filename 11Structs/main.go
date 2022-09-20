package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in Golang")

	chirag:= User{"chirag",25,"chirag@gmail.com"}

	fmt.Println(chirag)
	fmt.Printf("The details for chirag : %+v\n", chirag)
	fmt.Printf("The age of chirag is %v\n ", chirag.Age)
}

type User struct {
	Name string
	Age int
	Email string
}