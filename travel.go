package main

import (
	"math/rand"
	"time"
)

// Implements "Continue on trail" behaviour
func travel(rand rand.Rand, distance int, date time.Time, inv Inventory) (int, time.Time, Inventory) {
	return distance + 1, date.AddDate(0, 0, 1), inv
}
