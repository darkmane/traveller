package models


type BodyType int {
	Star = iota
	SmallGasGiant
	LargeGasGiant
	PlanetoidBelt
	RockyPlanet
}

type Body struct { }

func (b *Body) GetType() BodyType {
	return -1
}
