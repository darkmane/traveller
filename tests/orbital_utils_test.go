package tests

import (
	"fmt"
	"testing"

	"darkmane/traveller/models"
	"darkmane/traveller/util"
)

func TestGetSatelliteOrbits(t *testing.T) {
	orbits := util.GetSatelliteOrbits()

	for _, orbitKey := range []string{"ring", "close", "far", "extreme"} {
		if _, ok := orbits[orbitKey]; !ok {
			t.Errorf("%s missing from satellite orbits", orbitKey)
		}
	}
}

func TestGetAllOrbits(t *testing.T) {
	orbits := util.GetAllOrbits()
	for _, stellarSize := range models.GetAllStellarSizes() {
		for _, stellarClass := range models.GetAllStellarClasses() {
			simpleClass := fmt.Sprintf("%s0", stellarClass.ToString())
			if v, ok1 := orbits[stellarSize.ToString()]; ok1 {
				if _, ok2 := v[simpleClass]; !ok2 {
					t.Errorf("%s%s is missing from Stellar Orbits", stellarSize.ToString(), simpleClass)
				}
			} else {
				t.Errorf("%s is missing from Stellar Orbits", stellarSize.ToString())
			}
		}
	}
}
