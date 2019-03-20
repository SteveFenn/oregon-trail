package main

import (
	"math/rand"
)

// Parts ...
type Parts struct {
	wheel int
	oxen  int
	axel  int
}

// Inventory ...
type Inventory struct {
	money int
	oxen  int
	food  int
	parts Parts
}

// Implements "Continue on trail" behaviour
func travel(rand rand.Rand, distance int, date int, inv Inventory) (int, int, Inventory) {
	return distance + 1, date + 1, inv
}
