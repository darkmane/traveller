package models

import (
	"bytes"
	"encoding/json"
)

type BodyType int

const (
	RockyPlanet BodyType = iota
	SmallGasGiant
	LargeGasGiant
	PlanetoidBelt
	StellarBody
	Empty
)

type Body interface {
	GetType() BodyType
}

var bodyTypeToString = map[BodyType]string{
	StellarBody: "STAR",
	SmallGasGiant:  "SMALL_GAS_GIANT",
	LargeGasGiant:  "LARGE_GAS_GIANT",
	PlanetoidBelt: "PLANETOID_BELT",
	RockyPlanet: "ROCKY_PLANET",
}

var bodyTypeToID = map[string]BodyType{
	"STAR": StellarBody,
	"SMALL_GAS_GIANT": SmallGasGiant,
	"LARGE_GAS_GIANT": LargeGasGiant, 
	"PLANETOID_BELT": PlanetoidBelt,
	"ROCKY_PLANET": RockyPlanet,
}

// MarshalJSON marshals the enum as a quoted json string
func (bt BodyType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(bodyTypeToString[bt])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (bt *BodyType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'RockyPlanet' in this case.
	*bt = bodyTypeToID[j]
	return nil
}