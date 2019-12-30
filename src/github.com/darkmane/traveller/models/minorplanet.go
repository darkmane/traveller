package models

import "fmt"


type MinorPlanet struct {
   Planet
	 MainWorldId int
	 TradeClassifications []TradeClassifications
   Zone Zone
}
