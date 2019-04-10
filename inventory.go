package main

// Parts ...
type Parts struct {
	wheel int
	axel  int
}

// Inventory ...
type Inventory struct {
	money int
	oxen  int
	food  int
	parts Parts
}

func initInventory() Inventory {
	inventory := Inventory{money: 20, oxen: 4, food: 10, parts: Parts{wheel: 4, axel: 2}}
	return inventory
}
