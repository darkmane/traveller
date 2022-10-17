package models

import (
	"bytes"
	"encoding/json"
)

// TravelZone  Interstellar Travel Zone
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
