package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTravel(t *testing.T) {
	r := rand.NewSource(99)
	inv := initInventory()
	// firstFort(inv)
	var distance = 0
	var date, _ = time.Parse("20060102", "18900401")
	for i := 0; i < 10; i++ {
		distance, date, inv = travel(*rand.New(r), distance, date, inv)
		fmt.Printf("Distance: %d Date: %v Inventory: %v\n", distance, date.Format("2006-01-02"), inv)
	}
}
