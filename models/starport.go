package models

import (
	"bytes"
	"encoding/json"
)

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
