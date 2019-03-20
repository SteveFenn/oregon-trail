package main

import (
	"bufio"
	"fmt"
	"os"
)

// PartyMembers ...
type PartyMembers struct {
	name   string
	status string
}

// Party ...
func Party() []PartyMembers {
	fmt.Println("The Oregon Trail")
	fmt.Println("---------------------")
	leaderName := getUserInput("What is the name of your leader?:")
	member1 := getUserInput("What is the name of the second member?:")
	member2 := getUserInput("What is the name of the third member?:")
	member3 := getUserInput("What is the name of the fourth member?:")
	member4 := getUserInput("What is the name of the fifth member?:")

	status := [3]string{"well", "ill", "dead"}

	partyMembers := []PartyMembers{
		{
			leaderName,
			status[0],
		},
		{
			member1,
			status[0],
		},
		{
			member2,
			status[0],
		},
		{
			member3,
			status[0],
		},
		{
			member4,
			status[0],
		},
	}
	fmt.Println(partyMembers)
	return partyMembers
}

func getUserInput(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(question)
	text, _ := reader.ReadString('\n')
	return text
}
