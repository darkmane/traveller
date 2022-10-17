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
	StellarBodyString   = "star"
	SmallGasGiantString = "small_gas_giant"
	LargeGasGiantString = "large_gas_giant"
	PlanetoidBeltString = "planetoid_belt"
	RockyPlanetString   = "rocky_planet"
)

type Body interface {
	GetType() BodyType
}

var bodyTypeToString = map[BodyType]string{
	StellarBody:   StellarBodyString,
	SmallGasGiant: SmallGasGiantString,
	LargeGasGiant: LargeGasGiantString,
	PlanetoidBelt: PlanetoidBeltString,
	RockyPlanet:   RockyPlanetString,
}

var bodyTypeToID = map[string]BodyType{
	StellarBodyString:   StellarBody,
	SmallGasGiantString: SmallGasGiant,
	LargeGasGiantString: LargeGasGiant,
	PlanetoidBeltString: PlanetoidBelt,
	RockyPlanetString:   RockyPlanet,
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
