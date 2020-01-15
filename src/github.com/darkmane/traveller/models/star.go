package models

import (
	"math/bits"
	"bytes"
	"encoding/json"
)

const NEAR_ORBIT = -1
const FAR_ORBIT = 1<<(bits.UintSize-1) - 1 // MAX_INT
const CENTER = -1 << (bits.UintSize - 1)   // MIN_INT
type Star struct {
	Id    int
	Class StellarClass `json:"class"`
	Size  StellarSize  `json:"size"`
	Orbit int          `json:"orbit"`
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
