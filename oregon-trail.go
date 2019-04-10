package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Oregon Trail")
	Party()
	inventory := initInventory()
	firstFort(inventory)
}
