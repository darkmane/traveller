package models

import (
	"encoding/json"
)
const (
 body_type = "type"
 stellar = "stellar_orbit"
 planetary = "planetary_orbit"
)
type Orbit interface {
	GetType() BodyType
	GetOrbit() (int, int)
}

// type Orbit struct {
// 	Id             int
// 	StarSystemId   int
// 	StellarOrbit   int
// 	PlanetaryOrbit int
// 	bodyId         int64
// 	bodyType       BodyType
// }

// func (o *Orbit) GetType() BodyType {
// 	return o.bodyType
// }

type EmptyOrbit struct {
	StarSystemId   int
	StellarOrbit   int
	PlanetaryOrbit int
}

func (eo *EmptyOrbit) GetType() BodyType {
	return Empty
}

func (eo *EmptyOrbit) GetOrbit() (int, int) {
	return eo.StellarOrbit, eo.PlanetaryOrbit
}

func (eo *EmptyOrbit) MarshalJSON() ([]byte, error) {
	return json.Marshal(eo.ToMap())
}

func (eo *EmptyOrbit) ToMap() map[string]interface{} {
   m := make(map[string]interface{})
	 m[body_type] = "empty_orbit"
	 m[stellar] = eo.StellarOrbit
	 m[planetary] = eo.PlanetaryOrbit

	 return m
}
