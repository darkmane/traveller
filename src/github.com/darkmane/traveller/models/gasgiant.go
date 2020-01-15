package models

import (
	"errors"
	. "github.com/darkmane/traveller/util"
)

type GasGiant struct {
	Type BodyType
	Moons map[int]Planet
}

func (gg *GasGiant) GetType() BodyType {
	return gg.Type
}

func (gg *GasGiant) SetType(bt BodyType) error {

	if bt == LargeGasGiant || bt == SmallGasGiant {
		gg.Type = bt
		return nil
	}
	return errors.New("Incorrect BodyType")
}

func CreateGasGiant(dg *DiceGenerator) *GasGiant {
	gg := new(GasGiant)
	gg.SetType(LargeGasGiant)
	if (dg.Roll() % 2) == 0 {
		gg.SetType(SmallGasGiant)
	}

	gg.createSatellites(dg)

	return gg
}

func (gg *GasGiant)createSatellites(dg *DiceGenerator) {
	mod := 0
	moons := make(map[int]*Planet)
	if gg.Type == SmallGasGiant {
		mod = -4
	}
	numMoons := dg.RollDiceWithModifier(2, mod)
	ss := gg.getStarSystem()
	z := gg.getOrbitalZone()
	for counter := 0; counter < numMoons; counter++ {
		moon := createSatellite(dg, z, ss)
		orbit := LookUpOrbit(dg, (moon.Size == 0))
		moons[orbit] = moon
	}

}


func createSatellite(dg *DiceGenerator, zone Zone, ss *StarSystem) *Planet {
	return new(Planet)
}

func (gg *GasGiant) getStarSystem() *StarSystem {
	return new(StarSystem)
}

func (gg *GasGiant) getOrbitalZone() Zone {
	return INNER
}