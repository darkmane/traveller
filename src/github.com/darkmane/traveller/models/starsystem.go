package models

import (
	"fmt"
	"bytes"
	"encoding/json"	
	. "github.com/darkmane/traveller/util"
	"log"
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
	ScoutBase  bool       `json:"scout"`
	NavalBase  bool       `json:"naval"`
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

func (ss *StarSystem) FromMap(init map[string]interface{}) {
	for k, v := range init {
		switch k {
		case "x":
			ss.X = Interface2Int(v)
			break
		case "y":
			ss.Y = Interface2Int(v)
			break
		case "sector":
			ss.Sector = v.(string)
			break
		case "subsector":
			ss.SubSector = v.(string)
			break
		case "travel_zone":
			s := []byte(v.(string))
			ss.TravelZone.UnmarshalJSON(s)
			break
		case "scout":
			ss.ScoutBase = v.(bool)
			break
		case "navy":
			ss.NavalBase = v.(bool)
			break
		}
	}

	p := ss.Planet
	p.FromMap(init)
	ss.Planet = p
}

func (ss *StarSystem) ToMap() map[string]interface{} {
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
	os := make([]Orbit, 0)
	if ss.Orbits != nil { 
		log.Printf("Orbits is not nil: %v", ss.Orbits)
		os = ss.Orbits
	 }
	
	output["orbits"] = os
	
	return output
}

func (ss *StarSystem) UnmarshalJSON(b []byte) error{
	working_copy := make(map[string]interface{})
	err := json.Unmarshal(b, &working_copy)
	if err != nil { return err }
	ss.FromMap(working_copy)

	return nil
}

func (ss *StarSystem) MarshalJSON() ([]byte, error){
	return json.Marshal(ss.ToMap())
}
