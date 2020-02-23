package models

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
