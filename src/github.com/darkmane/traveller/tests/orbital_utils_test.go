package util

import (
	"fmt"
	"testing"

	"github.com/darkmane/traveller/models"
	"github.com/darkmane/traveller/util"
)

func TestGetSatelliteOrbits(t *testing.T) {
	orbits := util.GetSatelliteOrbits()

	for _, orbit_key := range []string{"ring", "close", "far", "extreme"} {
		if _, ok := orbits[orbit_key]; !ok {
			t.Errorf("%s missing from satellite orbits", orbit_key)
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
