package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Welcome to User Input"
	fmt.Println(welcome);

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating here")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating" ,input)

	fmt.Printf("the type of input is %T" , input)

	//comma ok syntax - error ok syntax


}