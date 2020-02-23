package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	. "github.com/darkmane/traveller/util"
)

type TravelZone int

const (
	Green TravelZone = iota
	Yellow
	Red
)

var travelZoneToString = map[TravelZone]string{
	Green:  "GREEN",
	Yellow: "YELLOW",
	Red:    "RED",
}

var travelZoneToID = map[string]TravelZone{
	"GREEN":  Green,
	"YELLOW": Yellow,
	"RED":    Red,
}

func (tz TravelZone) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(travelZoneToString[tz])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (tz *TravelZone) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'B' in this case.
	*tz = travelZoneToID[j]
	return nil
}

type Zone int

const (
	UNAVAILABLE Zone = iota
	INNER
	HABITABLE
	OUTER
)

var zoneToString = map[Zone]string{
	UNAVAILABLE: "UNAVAILABLE",
	INNER:       "INNER",
	HABITABLE:   "HABITABLE",
	OUTER:       "OUTER",
}

var zoneToID = map[string]Zone{
	"UNAVAILABLE": UNAVAILABLE,
	"INNER":       INNER,
	"HABITABLE":   HABITABLE,
	"OUTER":       OUTER,
}

func (sz Zone) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(zoneToString[sz])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (z *Zone) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'B' in this case.
	*z = zoneToID[j]
	return nil
}

func (z Zone) ToString() string {
	return zoneToString[z]
}

type Starport int

const (
	StarportA = iota
	StarportB
	StarportC
	StarportD
	StarportE
	StarportX
	StarportY
	StarportH
	StarportG
	StarportF
	StarportNone
)

var starportToID = map[string]Starport{
	"Class A Starport":       StarportA,
	"Class B Starport":       StarportB,
	"Class C Starport":       StarportC,
	"Class D Starport":       StarportD,
	"Class E Starport":       StarportE,
	"Class X Starport":       StarportX,
	"No Spaceport":           StarportY,
	"Primitive Spaceport":    StarportH,
	"Poor quality Spaceport": StarportG,
	"Good quality Spaceport": StarportF,
	"No Starport":            StarportNone,
}

var starportIDToString = map[Starport]string{
	StarportA:    "Class A Starport",
	StarportB:    "Class B Starport",
	StarportC:    "Class C Starport",
	StarportD:    "Class D Starport",
	StarportE:    "Class E Starport",
	StarportX:    "Class X Starport",
	StarportY:    "No Spaceport",
	StarportH:    "Primitive Spaceport",
	StarportG:    "Poor quality Spaceport",
	StarportF:    "Good quality Spaceport",
	StarportNone: "No Starport",
}

func (sp Starport) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(starportIDToString[sp])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (sp Starport) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'B' in this case.
	sp = starportToID[j]
	return nil
}

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

func NewStarSystem(initial map[string]interface{}, dg *DiceGenerator) *StarSystem {

	ss := new(StarSystem)

	p := NewPlanet(initial, dg)
	filled := ss.load(initial)
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

	ss.maxOrbits = MaxInt(orbitRoll, 1)
	ss.Orbits = make([]Orbit, ss.maxOrbits)
	log.Printf("At most %d are available", ss.maxOrbits)
	for k, _ := range ss.generateEmptyOrbits(dg) {
		log.Printf("Orbit %d of %d is empty", k, len(ss.Orbits))
		ss.Orbits[k-1] = &(EmptyOrbit{StarSystemId: ss.Id, StellarOrbit: k})
	}

	gasGiantOrbits := ss.generateGasGiants(dg)
	ss.Orbits = append(ss.Orbits, gasGiantOrbits...)

	for o, _ := range ss.getAvailableOrbits() {
		log.Printf("{0} is available", o)
	}

	log.Printf("New Star System : %v", filled)
	return ss
}

func (ss *StarSystem) load(init map[string]interface{}) map[string]bool {
	rv := make(map[string]bool)
	for k, v := range init {
		switch k {
		case x:
			ss.X = Interface2Int(v)
			rv[k] = true
			break
		case y:
			ss.Y = Interface2Int(v)
			rv[k] = true
			break
		case sector:
			ss.Sector = v.(string)
			rv[k] = true
			break
		case subsector:
			ss.SubSector = v.(string)
			rv[k] = true
			break
		case travel_zone:
			s := []byte(v.(string))
			ss.TravelZone.UnmarshalJSON(s)
			rv[k] = true
			break
		case scout:
			ss.ScoutBase = v.(bool)
			rv[k] = true
			break
		case naval:
			ss.NavalBase = v.(bool)
			rv[k] = true
			break
		}
	}
	p := ss.Planet
	p.FromMap(init)
	ss.Planet = p

	return rv
}

func (ss *StarSystem) FromMap(init map[string]interface{}) {
	ss.load(init)
}

func (ss *StarSystem) ToMap() map[string]interface{} {
	p := ss.Planet
	output := p.ToMap()
	dg := NewDiceGenerator("foo")
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
		log.Printf("Orbits is not nil: %v", ss.Orbits)
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

func (ss *StarSystem) generateEmptyOrbits(dg *DiceGenerator) map[int]interface{} {
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
	numberEmpty = MinInt(numberEmpty, ss.maxOrbits-1)
	for len(empty) < numberEmpty {
		empty[MinInt(dg.Roll(), ss.maxOrbits-1)] = true
	}

	return empty

}

func (ss *StarSystem) generateGasGiants(dg *DiceGenerator) []Orbit {
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
		numGG = MinInt(numGG, ss.maxOrbits)

		i := 0
		for i < numGG {
			log.Printf("Placing %d GG", i)
			i += 1
		}
	}
	return gasGiants
}

func (ss *StarSystem) placePlanetoidBelts(dg *DiceGenerator, gasGiantCount int) []Orbit {
	var pb []Orbit
	rollForPlanetoids := MaxInt(0, dg.Roll()-gasGiantCount)

	if rollForPlanetoids > 7 {
		return pb
	}

	log.Printf("Planetoid belts present")
	pbCount := 1
	switch MaxInt(0, dg.Roll()-gasGiantCount) {
	case 0:
		pbCount = 3
	case 1, 2, 3, 4, 5, 6:
		pbCount = 3
	}
	pbCount = MinInt(pbCount, len(ss.getAvailableOrbits()))
	var counter int
	counter = 0
	for counter < pbCount {
		log.Printf("Placing Planetoid Belt %d", counter)
		counter += 1
	}
	return pb
}

func (ss *StarSystem) getAvailableOrbits(zones ...Zone) map[int]interface{} {
	availableOrbits := ss.Stars[PRIMARY].GetOrbits(zones...)
	log.Printf("Available Orbits: %d", len(ss.Orbits))
	for i, o := range ss.Orbits {
		if o != nil {
			delete(availableOrbits, i)
			log.Printf("Orbit is nil")
		}
	}

	return availableOrbits
}
