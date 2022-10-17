package models

import (
	"encoding/json"
	"errors"

	"github.com/darkmane/traveller/util"
)

const (
	moons string = "moons"
)

type GasGiant struct {
	Type         BodyType
	Moons        map[int]Planet
	stellarOrbit int
}

func (gg *GasGiant) GetType() BodyType {
	return gg.Type
}

func (gg *GasGiant) GetOrbit() (int, int) {
	return gg.stellarOrbit, 0
}

func (gg *GasGiant) SetType(bt BodyType) error {

	if bt == LargeGasGiant || bt == SmallGasGiant {
		gg.Type = bt
		return nil
	}
	return errors.New("Incorrect BodyType")
}

func CreateGasGiant(dg *util.DiceGenerator) *GasGiant {
	gg := new(GasGiant)
	gg.SetType(LargeGasGiant)
	if (dg.Roll() % 2) == 0 {
		gg.SetType(SmallGasGiant)
	}

	gg.createSatellites(dg)

	return gg
}

func (gg *GasGiant) createSatellites(dg *util.DiceGenerator) {
	mod := 0
	moons := make(map[int]Planet)
	if gg.Type == SmallGasGiant {
		mod = -4
	}
	numMoons := dg.RollDiceWithModifier(2, mod)
	ss := gg.getStarSystem()
	z := gg.getOrbitalZone()
	for counter := 0; counter < numMoons; counter++ {
		moon := createSatellite(dg, z, ss)
		orbit := util.LookUpOrbit(dg, (moon.Size == 0))
		moons[orbit] = *moon
	}
	gg.Moons = moons
}

func createSatellite(dg *util.DiceGenerator, zone Zone, ss *StarSystem) *Planet {
	return new(Planet)
}

func (gg *GasGiant) getStarSystem() *StarSystem {
	return new(StarSystem)
}

func (gg *GasGiant) getOrbitalZone() Zone {
	return OUTER
}

func (gg *GasGiant) MarshalJSON() ([]byte, error) {
	return json.Marshal(gg.ToMap())
}

func (gg *GasGiant) UnmarshalJSON(b []byte) error {
	working_copy := make(map[string]interface{})
	err := json.Unmarshal(b, &working_copy)
	if err != nil {
		return err
	}
	gg.FromMap(working_copy)

	return nil
}

func (gg *GasGiant) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m[body_type] = gg.Type
	m[stellar] = gg.stellarOrbit
	m[moons] = gg.Moons

	return m
}

func (gg *GasGiant) FromMap(init map[string]interface{}) {
	for k, v := range init {
		switch k {
		case body_type:
			str := []byte(v.(string))
			gg.Type.UnmarshalJSON(str)
			break
		case stellar:
			gg.stellarOrbit = util.Interface2Int(v)
			break
		case moons:
			str := []byte(v.(string))
			json.Unmarshal(str, gg.Moons)
			break
		}
	}
}
