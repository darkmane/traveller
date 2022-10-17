package models

import (
	"bytes"
	"encoding/json"
)

// Zone stores orbital zone types
type Zone int

// Orbit Zones (UNAVAILABLE = inside star, INNER = inside habitable zone, OUTER = outside habitable zone )
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
