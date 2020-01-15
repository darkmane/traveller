package models

type Orbit struct {
	Id             int
	StarSystemId   int
	StellarOrbit   int
	PlanetaryOrbit int
	bodyId         int64
	bodyType       BodyType
}

func (o *Orbit) GetBody() (*Body, error) {
	return nil, nil
}

func (o *Orbit) SetBody(b *Body) error {
	o.bodyId = b.Id
	o.bodyType = b.GetType()
	return nil
}
