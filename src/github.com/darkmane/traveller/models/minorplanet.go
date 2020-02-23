package models

type MinorPlanet struct {
	Planet
	MainWorldId          int
	Classifications 	TradeClassifications
	Zone                 Zone
}

func (mp *MinorPlanet)GetType() BodyType {
	return mp.Planet.GetType()
}
