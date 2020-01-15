package models

type MinorPlanet struct {
	Planet
	MainWorldId          int
	Classifications 	TradeClassifications
	Zone                 Zone
}
