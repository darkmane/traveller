package util

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var stellarOrbits map[string]map[string]map[string][]int
var satelliteOrbits map[string]map[string]int

// GetAllOrbits Stellar orbits
func GetAllOrbits() map[string]map[string]map[string][]int {
	if stellarOrbits == nil {
		stellarOrbits = make(map[string]map[string]map[string][]int)
		rawStellarOrbits, exists := Get("/orbital/zones.yaml")
		if !exists {
			log.Printf("Stellar Orbital Zones missing")
		}
		err := yaml.Unmarshal(rawStellarOrbits, stellarOrbits)
		if err != nil {
			log.Printf("Unable to unmarshal Stellar Orbital zones: %v", err)
		}
	}

	return stellarOrbits
}

// GetSatelliteOrbits Get the satellite orbits tables
func GetSatelliteOrbits() map[string]map[string]int {
	if satelliteOrbits == nil {
		satelliteOrbits = make(map[string]map[string]int)
		rawSatelliteOrbits, exists := Get("/orbital/satellite.yaml")
		if exists {
			log.Printf("Satellite Orbital distances missing")
		}
		err := yaml.Unmarshal(rawSatelliteOrbits, satelliteOrbits)
		if err != nil {
			log.Printf("Unable to unmarshal satellite orbits: %v", err)
		}
	}

	return satelliteOrbits
}

// LookUpOrbit Lookup the satellite orbits
func LookUpOrbit(dg *DiceGenerator, ring bool) int {
	satelliteOrbits := GetSatelliteOrbits()
	key := "extreme"
	var roll int
	roll = dg.Roll()
	switch {
	case roll <= 7:
		key = "close"
	case roll <= 11:
		key = "far"
	}

	var rv int
	if options, ok := satelliteOrbits[key]; ok {
		roll = dg.Roll()
		if len(options) == 6 {
			roll = dg.RollDiceWithModifier(1, 0)
		}
		rv = options[fmt.Sprintf("%d", roll)]
	} else {
		ProcessError(fmt.Errorf("Unknown Key for Satellite Orbits [%v]", key))
	}
	return rv
}
