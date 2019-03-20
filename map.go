package main

// LocationType ...
type LocationType int

const (
	// Fort ...
	Fort LocationType = iota + 1
	// River ...
	River LocationType = iota + 1
	// Forest ...
	Forest LocationType = iota + 1
)

// Location ...
type Location struct {
	name         string
	locationType LocationType
	distance     int
}

var locations = [...]Location{Location{"Independence", Fort, 0}}
