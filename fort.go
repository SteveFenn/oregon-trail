package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func notMain() {
	inv := Inventory{money: 20, oxen: 4, food: 10, parts: Parts{wheel: 4, axel: 2}}
	makeDecision(inv)
}

func makeDecision(inv Inventory) {
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
		shop(inv)
	case "3":
		fmt.Println("Good choice")
	case "4":
		fmt.Println("Good choice")
	default:
		fmt.Println("Enter correct input")
		makeDecision(inv)
	}
}

func continueTrail() {
	fmt.Println("Back on the trail...")
}

func shop(inv Inventory) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("You have money")
	fmt.Println("What would you like to buy?")
	fmt.Println("1. Oxen")
	fmt.Println("2. Food")
	fmt.Println("3. Spare parts")

	text, _ := reader.ReadString('\n')
	formattedText := strings.TrimSpace(text)

	switch formattedText {
	case "1":
		buyOxen(inv)
	case "2":
		fmt.Println("Good choice")
	case "3":
		fmt.Println("Good choice")
	default:
		fmt.Println("Enter correct input")
		shop(inv)
	}
}

func buyOxen(inv Inventory) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("You have money")
	fmt.Println("How many Oxen would you like to buy?")

	text, _ := reader.ReadString('\n')
	newOxen, _ := strconv.Atoi(strings.TrimSpace(text))
	inv.oxen = inv.oxen + newOxen

	fmt.Printf("You have purchased %d Oxen\n", newOxen)
	fmt.Printf("You now have %d Oxen\n", inv.oxen)
}

func buyFood() {
	// reader := bufio.NewReader(os.Stdin)
	fmt.Println("You have money")
	fmt.Println("How much Food would you like to buy?")
	// text, _ := reader.ReadString('\n')

}
