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

// Status ...
type Status int

// Dead Ill Well ...
const (
	Dead Status = 0
	Ill  Status = 1
	Well Status = 2
)

// Party ...
func Party() []PartyMembers {
	fmt.Println("The Oregon Trail")
	fmt.Println("---------------------")
	leaderName := getUserInput("What is the name of your leader?:")
	member1 := getUserInput("What is the name of the second member?:")
	member2 := getUserInput("What is the name of the third member?:")
	member3 := getUserInput("What is the name of the fourth member?:")
	member4 := getUserInput("What is the name of the fifth member?:")

	partyMembers := []PartyMembers{
		{
			leaderName,
			Well,
		},
		{
			member1,
			Well,
		},
		{
			member2,
			Well,
		},
		{
			member3,
			Well,
		},
		{
			member4,
			Well,
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
