package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestTravel(t *testing.T) {
	r := rand.NewSource(99)
	var inv Inventory
	var distance = 0
	var date = 0
	for i := 0; i < 10; i++ {
		distance, date, inv = travel(*rand.New(r), distance, date, inv)
		fmt.Printf("Distance: %d Date: %d Inventory: %v\n", distance, date, inv)
	}
}
