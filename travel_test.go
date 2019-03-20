package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestTravel(t *testing.T) {
	r := rand.NewSource(99)
	var inv inventory
	for i := 0; i < 10; i++ {
		inv = travel(*rand.New(r), inv)
		fmt.Println(inv)
	}

}
