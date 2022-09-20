package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// fmt.Println("Welcome to if else");

	// var loginCount int = 10

	// var message string;

	// if loginCount > 10 {
	// 	message = "bda hai"
	// } else if loginCount < 10 {
	// 	message = "luss"
	// } else {
	// 	message = "decade"
	// }

	// fmt.Println("The result is", message)

	fmt.Println("Welcome to switch case");

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6)+1
	fmt.Println("The dice number is", diceNumber) 

	switch diceNumber {
	case 1:
		fmt.Println("The number is 1")
	case 2:
		fmt.Println("The number is 2")
	case 3:
		fmt.Println("The number is 3")
	case 4:
		fmt.Println("The number is 4")
	case 5:
		fmt.Println("The number is 5")
	case 6:
		fmt.Println("The number is 6")
	default:
		fmt.Println("What is that?")

	}

}

