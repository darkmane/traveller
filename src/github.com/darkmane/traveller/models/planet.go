package models

import "fmt"

const (
    maxLand = 0.2;
    minLand = 0.2;
    water = 0.02;
    maxIce = 0.85;
    maxCloud = 0.8;
    minIce = 0.55;
    minCloud = 0.4;

    ring = 0;
    small = -1;
)
type Planet struct {
	Id int
	UniversalPlanetProfile
  maxOrbits int = 0
	Satellites map[int]Planet
	
}
