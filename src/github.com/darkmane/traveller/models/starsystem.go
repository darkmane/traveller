package models

import (
	"fmt"
	"bytes"
	"encoding/json"
	. "github.com/darkmane/traveller/util"
)

type TravelZone int

const (
	Green TravelZone = iota 
	Yellow
	Red
)
var travelZoneToString = map[TravelZone]string {
	Green: "GREEN",
	Yellow: "YELLOW",
	Red: "RED",
}

var travelZoneToID = map[string]TravelZone {
	"GREEN": Green,
	"YELLOW": Yellow,
	"RED": Red,
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

var zoneToString = map[Zone]string {
	UNAVAILABLE: "UNAVAILABLE",
	INNER: "INNER",
	HABITABLE: "HABITABLE",
	OUTER: "OUTER",
}

var zoneToID = map[string]Zone {
	"UNAVAILABLE": UNAVAILABLE,
	"INNER": INNER,
	"HABITABLE": HABITABLE,
	"OUTER": OUTER,
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
	"Class A Starport": StarportA,
    "Class B Starport": StarportB,
    "Class C Starport": StarportC,
    "Class D Starport": StarportD,
    "Class E Starport": StarportE,
    "Class X Starport": StarportX,
    "No Spaceport": StarportY,
    "Primitive Spaceport": StarportH,
    "Poor quality Spaceport": StarportG,
    "Good quality Spaceport": StarportF,
    "No Starport": StarportNone,
}

var starportIDToString = map[Starport]string {
	StarportA: "Class A Starport",
    StarportB: "Class B Starport",
    StarportC: "Class C Starport",
    StarportD: "Class D Starport",
    StarportE: "Class E Starport",
    StarportX: "Class X Starport",
    StarportY: "No Spaceport",
    StarportH: "Primitive Spaceport",
    StarportG: "Poor quality Spaceport",
    StarportF: "Good quality Spaceport",
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
	Planet
	X          int        `json:"x"`
	Y          int        `json:"y"`
	Sector     string     `json:"sector"`
	SubSector  string     `json:"subsector"`
	TravelZone TravelZone `json:"travel_zone,string"`
	ScoutBase  bool       `json:"scout,string"`
	NavalBase  bool       `json:"naval,string"`
	Orbits     []Orbit    `json:"orbits"`
}

//
func (ss *StarSystem) Coordinate() string {
	return fmt.Sprint("%d-%d", ss.X, ss.Y)
}

func NewStarSystem(initial map[string]int) (*StarSystem){
	ss :=  new(StarSystem)

	return ss
}

func (ss *StarSystem) generate() {

}

func (ss *StarSystem) Load(init map[string]int) {
	if val, ok := init["x"]; ok {
		ss.X = val
	}
	if val, ok := init["y"]; ok {
		ss.Y = val
	}
	ss.generate()
	if val, ok := init["size"]; ok {
		ss.Size = val
	}
	if val, ok := init["atmosphere"]; ok {
		ss.Atmosphere = val
	}
	if val, ok := init["hydro"]; ok {
		ss.Hydro = val
	}
	if val, ok := init["population"]; ok {
		ss.Population = val
	}
	if val, ok := init["government"]; ok {
		ss.Government = val
	}
	if val, ok := init["lawlevel"]; ok {
		ss.LawLevel = val
	}
	if val, ok := init["techlevel"]; ok {
		ss.TechLevel = val
	}
}

func (ss *StarSystem) UnmarshalJSON(b []byte) error{
	return nil
}

func (ss *StarSystem) MarshalJSON() ([]byte, error){
	p := ss.Planet
	output := p.ToMap()
	dg := NewDiceGenerator("foo")
	ss.Classifications = calculateTradeFacilities(&dg, &p, ss, HABITABLE)
	
	output["x"] = ss.X
	output["y"] = ss.Y
	output["sector"] = ss.Sector
	output["subsector"] = ss.SubSector
	output["travel_zone"] = ss.TravelZone
	output["scout"] = ss.ScoutBase
	output["naval"] = ss.NavalBase
	output["orbits"] = ss.Orbits
	output["classifications"] = ss.Classifications
	return json.Marshal(output)
}
