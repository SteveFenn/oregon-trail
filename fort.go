package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	makeDecision()
}

func makeDecision() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the fort what would you like to do?")
	fmt.Println("1. Continue with the trail.")
	fmt.Println("2. Buy suppliers from shop.")
	fmt.Println("3. Rest.")
	fmt.Println("4. Attempt to Trade")

	text, _ := reader.ReadString('\n')
	formattedText := strings.TrimSpace(text)

	switch formattedText {
	case "1":
		continueTrail()
	case "2":
		fmt.Println("Good choice")
	case "3":
		fmt.Println("Good choice")
	case "4":
		fmt.Println("Good choice")
	default:
		fmt.Println("Enter correct input")
		makeDecision()
	}
}

func continueTrail() {
	fmt.Println("Back on the trail...")
}

func shop() {
	reader := bufio.NewReader(os.Stdin)
	fmt.println("You have money")
	fmt.Println("What would you like to buy?")
	fmt.Println("1. Oxen")
	fmt.Println("2. Food")
	fmt.Println("3. Spare parts")

	text, _ := reader.ReadString('\n')
	formattedText := strings.TrimSpace(text)

	switch formattedText {
	case "1":
		buyOxen()
	case "2":
		fmt.Println("Good choice")
	case "3":
		fmt.Println("Good choice")
	default:
		fmt.Println("Enter correct input")
		shop()
	}
}

func buyOxen() {
	reader := bufio.NewReader(os.Stdin)
	fmt.println("You have money")
	fmt.Println("How many Oxen would you like to buy?")

	text, _ := reader.ReadString('\n')
	formattedText := strings.TrimSpace(text)
}
