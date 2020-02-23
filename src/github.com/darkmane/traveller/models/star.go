package models

import (
	"math/bits"
	"bytes"
	"encoding/json"
	. "github.com/darkmane/traveller/util"
)

const NEAR_ORBIT = -1
const FAR_ORBIT = 1<<(bits.UintSize-1) - 1 // MAX_INT
const CENTER = -1 << (bits.UintSize - 1)   // MIN_INT
const(
    SOLITARY int = 7;
    BINARY int = 11;
	TRINARY int = 12;
)
const (
	class string = "class"
	size string = "size"
	orbit string = "orbit"
)
type Star struct {
	Class StellarClass `json:"class"`
	Size  StellarSize  `json:"size"`
	SizeFraction int   `json:"fraction"`
	Orbit int          `json:"orbit"`
}

func (s *Star)GetType() BodyType {
	return StellarBody
}

func (s *Star) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToMap())
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *Star) UnmarshalJSON(b []byte) error {
	working_copy := make(map[string]interface{})
	err := json.Unmarshal(b, &working_copy)
	if err != nil { return err }
	s.FromMap(working_copy)

	return nil
}

func (s *Star) FromMap(init map[string]interface{}) {
	for k, v := range init {
		switch k {
			case class:
				str := []byte(v.(string))
				s.Class.UnmarshalJSON(str)
				break
			case size:
				str := []byte(v.(string))
				s.Size.UnmarshalJSON(str)
				break
			case orbit:
				s.Orbit = Interface2Int(v)
				break
		}
	}
}

func (s *Star) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m[class] = s.Class
	m[size] = s.Size
	m[orbit] = s.Orbit

	return m
}

type StellarClass int

const (
	B StellarClass = iota
	A
	M
	K
	G
	F
)

var stellarClassToString = map[StellarClass]string{
	B: "B",
	A: "A",
	M: "M",
	K: "K",
	G: "G",
	F: "F",
}

var stellarClassToID = map[string]StellarClass{
	"B": B,
	"A": A,
	"M": M,
	"K": K,
	"G": G,
	"F": F,
}

func (sc StellarClass) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(stellarClassToString[sc])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (sc *StellarClass) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'B' in this case.
	*sc = stellarClassToID[j]
	return nil
}

func (sc *StellarClass) ToString() string {
	return stellarClassToString[*sc]
}

type StellarSize int

const (
	Ia StellarSize = iota
	Ib
	II
	III
	IV
	V
	VI
	D
)
var stellarSizeToString = map[StellarSize]string{
	Ia: "Ia",
	Ib: "Ib",
	II: "II",
	III: "III",
	IV: "IV",
	V: "V",
	VI: "VI",
	D: "D",
}

var stellarSizeToID = map[string]StellarSize{
	"Ia": Ia,
	"Ib": Ib,
	"II": II,
	"III": III,
	"IV": IV,
	"V": V,
	"VI": VI,
	"D": D,
}
func (ss StellarSize) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(stellarSizeToString[ss])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (ss *StellarSize) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'B' in this case.
	*ss = stellarSizeToID[j]
	return nil
}

func (ss *StellarSize) ToString() string {
	return stellarSizeToString[*ss]
}

type StarPosition int

const (
	PRIMARY StarPosition = iota
	SECONDARY
	TERTIARY
)

var starPositionToString = map[StarPosition]string{
	PRIMARY: "PRIMARY",
	SECONDARY: "SECONDARY",
	TERTIARY: "TERTIARY",
}

var starPositionToID = map[string]StarPosition{
	"PRIMARY": PRIMARY,
	"SECONDARY": SECONDARY,
	"TERTIARY": TERTIARY,
}

func (sp StarPosition) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(starPositionToString[sp])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (sp *StarPosition) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'B' in this case.
	*sp = starPositionToID[j]
	return nil
}

func (s *Star)GetOrbits(zones ...Zone) map[int]interface{} {

	var orbits []int
	
	orbit := GetAllOrbits()
	stellarOrbits := orbit[s.Size.ToString()][s.Class.ToString()]

	for z := range zones {
		orbits = append(orbits, stellarOrbits[Zone(z).ToString()]...)
	}

	var returnOrbits map[int]interface{}
	for o := range orbits {
		returnOrbits[o] = true
	}
	return returnOrbits
}

func generateStars(dg *DiceGenerator, population int, atmos int) map[StarPosition]*Star {
	rollClass := dg.Roll()
	rollSize := dg.Roll()
	star := dg.Roll()
	
	rv := make(map[StarPosition]*Star)

	if (population > 7 || (atmos > 3 && atmos < 10)) {
		rollClass += 4;
		rollSize += 4;
	}
	var sClass StellarClass = M
	var sSize StellarSize = V
	
	switch rollSize {
	case 0:
		sSize = Ia
	case 1:
		sSize = Ib
	case 2:
		sSize = II
	case 3:
		sSize = III
	case 4:
		sSize = IV
	case 11:
		sSize = VI
	case 12:
		sSize = D
	}

	switch rollClass {
	case 0, 1:
		sClass = B
		if sSize == VI {
			sSize = V
		}
	case 2:
		sClass = A
		if sSize == VI {
			sSize = V
		}
	case 3, 4, 5, 6, 7:
		sClass = M
		if sSize == IV {
			sSize = V
		}
	case 8:
		sClass = K
	case 9:
		sClass = G
	case 10, 11, 12:
		sClass = F
		if sSize == VI {
			sSize = V
		}
	}

	fraction := 0
	if dg.RollDice(1) > 3 {
		fraction = 5
	}

	rv[PRIMARY] = &Star{sClass, sSize, fraction, CENTER}
	if star > BINARY {
		rv[SECONDARY] = generateCompanionStar(dg, rollClass, rollSize, SECONDARY)
	}
	
	if star > TRINARY {
		rv[TERTIARY] = generateCompanionStar(dg, rollClass, rollSize, TERTIARY)
	}
	return rv
}

func generateCompanionStar(dg *DiceGenerator, classMod int, sizeRoll int, position StarPosition) *Star {

	var rollClass int = dg.Roll()
	var rollSize int = dg.Roll()
	var rollOrbit int = dg.Roll()
	var orbitRoll int = dg.Roll()

	orbit := CENTER
	sClass := M
	sSize := D

	rollClass += classMod
	rollSize += sizeRoll
	if position == TERTIARY {
		rollOrbit += 4
	}

	switch rollClass {
	case 1: 
		sClass = B
	case 2:
		sClass = A
	case 3, 4:
		sClass = F
	case 5,6:
		sClass = G
	case 7, 8:
		sClass = K
	}

	switch rollSize {
	case 0:
		sSize = Ia
	case 1:
		sSize = Ib
	case 2:
		sSize = II
	case 3:
		sSize = III
	case 4:
		sSize = IV
	case 5, 6:
		sSize = D
	case 7, 8:
		sSize = V
	case 9:
		sSize = VI
	}

	fraction := 0
	if dg.RollDice(1) > 3 {
		fraction = 5
	}

	switch {
	case rollOrbit < 4:
		orbit = NEAR_ORBIT
	case rollOrbit == 12:
		orbit = FAR_ORBIT
	default:
		orbit = rollOrbit - 3
		if rollOrbit > 6 {
			orbit += dg.Roll()
		}
	}

	switch sSize {
	case III:
		orbitRoll += 4
	case Ia, Ib, II:
		orbitRoll += 8
	}

	switch sClass {
	case M:
		orbitRoll -= 4
	case K:
		orbitRoll -= 2
	}

	return &Star{sClass, sSize, fraction, rollOrbit}

}


func (s *Star)GetNearestMajorClass() string {
	return "M5"
}