package main

import (
	"math/rand"
)

// Implements "Continue on trail" behaviour
func travel(rand rand.Rand, distance int, date int, inv Inventory) (int, int, Inventory) {
	return distance + 1, date + 1, inv
}
