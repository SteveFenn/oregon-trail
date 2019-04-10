package main

import (
	"math/rand"
	"time"
)

type event struct {
	title       string
	description string
	changes     effect
}

type effect struct {
	partyChange
	inventoryChange
	dateChange int
}

type inventoryChange struct {
	item   string
	change int
}
type partyChange struct {
	member int
	change Status
}

// Events ...
func Events() {
	// Create type of events
	// Check if you get an event
	// Pick an event from list
	// Update inventory

	// event := pickEvent()
	updateInventory()
}

func updateInventory() {

}

func pickEvent() event {
	events := []event{
		{
			"Oh noes! Bear Attack!!!",
			"One of your oxen has been eaten by a bear, you find his remains scattered across the camp.",
			effect{
				inventoryChange: inventoryChange{"oxen", -1},
			},
		},
		{
			"You're a hero!",
			"You saved the local sheriff's daughter from robbers. You've earned a reward!",
			effect{
				inventoryChange: inventoryChange{"money", 50},
			},
		},
		{
			"You've gone the wrong bloody way!",
			"You trusted your instincts and they were wrong! You've gone off track and it takes you a couple of days to get back.",
			effect{
				dateChange: 2,
			},
		},
		{
			"Hold me closer tiny dancer",
			"Your dancing impresses the local indians and they shower you with food. All those days infront of the mirror has paid off!",
			effect{
				inventoryChange: inventoryChange{"food", 20},
			},
		},
	}
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return events[r.Intn(len(events))]
}
