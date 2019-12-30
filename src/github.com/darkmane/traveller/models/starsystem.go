package models

import (
				"fmt"
				"encoding/json"
)

type TravelZone int {
	Green=iota
	Yellow
	Red
}

type Zone int {
	UNAVAILABLE=iota
	INNER
	HABITABLE
	OUTER
}

type StarSystem struct {
				X          int `json:"x"`
				Y          int `json:"y"`
				Name       string `json:"name"`
				Sector     string `json:"sector"`
				SubSector  string `json:"subsector"`
				TravelZone TravelZone `json:"travel_zone"`
				ScoutBase  bool `json:"scout"`
				NavalBase  bool `json:"naval"`
				Orbits     []Orbit `json:"orbits"`
}

//
func (ss *StarSystem) Coordinate() string {
	return fmt.Sprint("%d-%d", ss.X, ss.Y)
}
