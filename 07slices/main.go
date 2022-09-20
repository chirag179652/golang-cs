package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is the slice study")

	var playerList = []string{"Lisandro", "Malacia", "Bruno"}
	fmt.Println("The list of the players is ", playerList)

	fmt.Printf("The type of the list is %T \n", playerList)
	playerList = append(playerList, "Eriksen", "Antony")
	fmt.Println("The list of the players is ", playerList)

	var playerList2 = playerList[:3]
	fmt.Println("The list of the new  players is ", playerList2)

	goals := make([]int, 4)
	goals[0] = 52
	goals[1] = 25
	goals[2] = 18
	goals[3] = 22
	fmt.Println("The value for goals is", goals)
	goals = append(goals, 91, 31)
	fmt.Println("The value for goals is", goals)
	sort.Ints(goals)
	fmt.Println("The value for goals is", goals)

	//Remove a value from slices based on index

	var courses = []string{"javascript", "python", "ruby", "react", "node"}

	index := 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
