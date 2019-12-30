package models

type Orbit struct {
	Id int
	StarSystemId int
	StellarOrbit int
	PlanetaryOrbit int
	bodyId int

}

func (o *Orbit) GetBody() Body, Error {

} 

func (o *Orbit) SetBody(b *Body) Error {

}
