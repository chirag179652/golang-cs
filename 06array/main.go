package main

import "fmt"

func main() {
	fmt.Println("This is the array study")

	var playerList [3]string

	playerList[0] = "Licha"
	playerList[2] = "Antony"

	fmt.Println("The list for players is", playerList )
	fmt.Println("The list length for players is", len(playerList))

	var coachList = [3]string{"Ole", "ETH", "Jose" }

	fmt.Println("coach list is", coachList)


}