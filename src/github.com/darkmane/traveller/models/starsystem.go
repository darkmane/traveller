package models

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/darkmane/traveller/util"
)

type StarSystem struct {
	Stars map[StarPosition]*Star `json:"stars"`
	*Planet
	X          int        `json:"x"`
	Y          int        `json:"y"`
	Sector     string     `json:"sector"`
	SubSector  string     `json:"subsector"`
	TravelZone TravelZone `json:"travel_zone,string"`
	maxOrbits  int        `json:"-"`
	Orbits     []Orbit    `json:"orbits"`
}

const (
	x           string = "x"
	y           string = "y"
	sector      string = "sector"
	subsector   string = "subsector"
	travel_zone string = "travel_zone"
	scout       string = "scout"
	naval       string = "naval"
	orbits      string = "orbits"
)

//
func (ss *StarSystem) Coordinate() string {
	return fmt.Sprint("%d-%d", ss.X, ss.Y)
}

func NewStarSystem(initial map[string]interface{}, dg *util.DiceGenerator) *StarSystem {

	ss := new(StarSystem)

	p := NewPlanet(initial, dg)
	ss.FromMap(initial)
	ss.Planet = p
	ss.Stars = generateStars(dg, ss.Size, ss.Atmosphere)

	orbitRoll := dg.Roll()
	switch ss.Stars[PRIMARY].Size {
	case III:
		orbitRoll += 4
	case Ia, Ib, II:
		orbitRoll += 8
	}

	switch ss.Stars[PRIMARY].Class {
	case M:
		orbitRoll -= 4
	case K:
		orbitRoll -= 2
	}

	ss.maxOrbits = util.MaxInt(orbitRoll, 1)
	ss.Orbits = make([]Orbit, ss.maxOrbits+1)
	log.Trace().Int("maxOrbits", ss.maxOrbits).Msg(fmt.Sprintf("At most %d are available", ss.maxOrbits))
	for k := range ss.generateEmptyOrbits(dg) {
		log.Trace().Msg(fmt.Sprintf("Orbit %d of %d is empty", k, len(ss.Orbits)))
		ss.Orbits[k-1] = &(EmptyOrbit{StarSystemId: ss.Id, StellarOrbit: k})
	}

	gasGiantOrbits := ss.generateGasGiants(dg)
	for _, gg := range gasGiantOrbits {
		so, _ := gg.GetOrbit()
		ss.Orbits[so] = gg
	}

	planetoids := ss.placePlanetoidBelts(dg)
	for _, pb := range planetoids {
		so, _ := pb.GetOrbit()
		ss.Orbits[so] = pb
	}
	// ss.Orbits = append(ss.Orbits, planetoids...)
	ss.generateMinorPlanets(dg)

	return ss
}

func (ss *StarSystem) FromMap(init map[string]interface{}) {

	for k, v := range init {
		switch k {
		case x:
			ss.X = util.Interface2Int(v)
			break
		case y:
			ss.Y = util.Interface2Int(v)
			break
		case sector:
			ss.Sector = v.(string)
			break
		case subsector:
			ss.SubSector = v.(string)
			break
		case travel_zone:
			s := []byte(v.(string))
			ss.TravelZone.UnmarshalJSON(s)
			break
		case scout:
			ss.ScoutBase = v.(bool)
			break
		case naval:
			ss.NavalBase = v.(bool)
			break
		}
	}

	ss.Planet.FromMap(init)
}

func (ss *StarSystem) ToMap() map[string]interface{} {
	p := ss.Planet
	output := p.ToMap()
	dg := util.NewDiceGenerator("foo")
	ss.Classifications = calculateTradeFacilities(&dg, p, ss, HABITABLE)

	output[x] = ss.X
	output[y] = ss.Y
	output[sector] = ss.Sector
	output[subsector] = ss.SubSector
	output[travel_zone] = ss.TravelZone
	output[scout] = ss.ScoutBase
	output[naval] = ss.NavalBase
	os := make([]Orbit, 0)
	if ss.Orbits != nil {
		log.Trace().Interface("orbits", ss.Orbits).Msg(fmt.Sprintf("Orbits is not nil: %v", ss.Orbits))
		os = ss.Orbits
	}

	output[orbits] = os

	return output
}

func (ss *StarSystem) UnmarshalJSON(b []byte) error {
	working_copy := make(map[string]interface{})
	err := json.Unmarshal(b, &working_copy)
	if err != nil {
		return err
	}
	ss.FromMap(working_copy)

	return nil
}

func (ss *StarSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(ss.ToMap())
}

func (ss *StarSystem) GetBodies(bodyTypes ...BodyType) []Orbit {
	var bodies []Orbit
	bodies = make([]Orbit, 0)
	for _, bt := range bodyTypes {
		for _, o := range ss.Orbits {
			if o != nil && o.GetType() == bt {
				bodies = append(bodies, o)
			}
		}
	}
	return bodies
}

func (ss *StarSystem) generateEmptyOrbits(dg *util.DiceGenerator) map[int]interface{} {
	empty := make(map[int]interface{})
	emptyRoll := dg.RollDice(1)
	numberEmptyRoll := dg.RollDice(1)

	switch ss.Stars[PRIMARY].Class {
	case A, B:
		emptyRoll += 1
		numberEmptyRoll += 1
	}

	numberEmpty := 0

	if emptyRoll > 3 {
		switch numberEmptyRoll {
		case 1, 2:
			numberEmpty = 1
		case 3:
			numberEmpty = 2
		default:
			numberEmpty = 3
		}
	}
	numberEmpty = util.MinInt(numberEmpty, ss.maxOrbits-1)
	for len(empty) < numberEmpty {
		empty[util.MinInt(dg.Roll(), ss.maxOrbits-1)] = true
	}

	return empty

}

func (ss *StarSystem) generateGasGiants(dg *util.DiceGenerator) []Orbit {
	var gasGiants []Orbit
	if dg.Roll() < 10 {
		numGG := 0
		switch dg.Roll() {
		case 1, 2, 3:
			numGG = 1
		case 4, 5:
			numGG = 2
		case 6, 7:
			numGG = 3
		case 8, 9, 10:
			numGG = 4
		case 11, 12:
			numGG = 5
		}
		numGG = util.MinInt(numGG, ss.maxOrbits)

		i := 0
		for i < numGG {
			log.Trace().Msg(fmt.Sprintf("Placing %d GG", i))
			gg := CreateGasGiant(dg)
			gasGiants = append(gasGiants, gg)
			i += 1
		}
	}
	return gasGiants
}

func (ss *StarSystem) generateMinorPlanets(dg *util.DiceGenerator) {

	message := ""
	for z := range [3]Zone{INNER, HABITABLE, OUTER} {
		zone := Zone(z)
		message += fmt.Sprintf("%v: %d", zone.ToString(), len(ss.getAvailableOrbits(zone)))
		for mainOrbit := range ss.getAvailableOrbits(zone) {
			mp := newMinorPlanet(dg, *ss.Stars[PRIMARY], mainOrbit, zone, ss)
			ss.Orbits[mainOrbit] = mp
		}
	}
	log.Trace().Msg(message)
}

func (ss *StarSystem) placePlanetoidBelts(dg *util.DiceGenerator) []Orbit {
	var pb []Orbit

	gasGiantCount := len(ss.GetBodies(SmallGasGiant, LargeGasGiant))
	rollForPlanetoids := util.MaxInt(0, dg.Roll()-gasGiantCount)

	if rollForPlanetoids > 7 {
		return pb
	}

	log.Trace().Msg(fmt.Sprintf("Planetoid belts present"))
	pbCount := 1
	switch util.MaxInt(0, dg.Roll()-gasGiantCount) {
	case 0:
		pbCount = 3
	case 1, 2, 3, 4, 5, 6:
		pbCount = 3
	}
	pbCount = util.MinInt(pbCount, len(ss.getAvailableOrbits()))
	var counter int
	counter = 0
	for counter < pbCount {
		log.Trace().Msg(fmt.Sprintf("Placing Planetoid Belt %d", counter))
		counter += 1
	}
	return pb
}

func (ss *StarSystem) getAvailableOrbits(zones ...Zone) map[int]interface{} {
	availableOrbits := ss.Stars[PRIMARY].GetOrbits(zones...)

	for i, o := range ss.Orbits {
		if o != nil {
			delete(availableOrbits, i)
		}
	}

	for i := range availableOrbits {
		if i > ss.maxOrbits {
			delete(availableOrbits, i)
		}
	}

	return availableOrbits
}
