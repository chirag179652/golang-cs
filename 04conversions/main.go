package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza app");
	fmt.Println("Please rate our Pizza between 1 to 5");

	reader := bufio.NewReader(os.Stdin);
	input,_ := reader.ReadString('\n');

	fmt.Println("Thanks for rating",input);
	numRating , err := strconv.ParseFloat( strings.TrimSpace(input) , 64)
	
	if err!= nil {
		fmt.Println("There was an error", err)
	} else {
		fmt.Println("Your new rating is added by one and is" ,numRating+1)
	}

}