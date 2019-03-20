package main

import (
  "fmt"
)

type event struct {
    title string
    description string
    changes []eventChange
}

type eventChange struct {
  inventoryItem string
  change int
}

func main(){
  // Create type of events
  // Check if you get an event
  // Pick an event from list
  // Update inventory
  events := []event{
    {
      "Oh noes! Bear Attack!!!",
      "One of your oxen has been eaten by a bear, you find his remains scattered across the camp.",
      []eventChange{
        {"oxen", -1},
      },
    },
    {
      "You're a hero!",
      "You saved the local sheriff's daughter from robbers. You've earned a reward!",
      []eventChange{
        {"money", 50},
      },
    },
    {
      "You've gone the wrong bloody way!",
      "You trusted your instincts and they were wrong! You've gone off track and it takes you a couple of days to get back.",
      []eventChange{
        {"days", -2},
      },
    },
    {
      "Hold me closer tiny dancer",
      "Your dancing impresses the local indians and they shower you with food. All those days infront of the mirror has paid off!",
      []eventChange{
        {"food", 20},
      },
    },
  }
  fmt.Println(events)
}
