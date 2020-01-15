package models

import (
	. "github.com/darkmane/traveller/util"
)

const (
	maxLand  = 0.2
	minLand  = 0.2
	water    = 0.02
	maxIce   = 0.85
	maxCloud = 0.8
	minIce   = 0.55
	minCloud = 0.4

	Ring  = 0
	Small = -1
)

type Planet struct {
	Id int `json:"-"`
	Name string `json:"name"`
	UniversalPlanetProfile
	Port Starport `json:"starport,string"`
	maxOrbits  int `json:"-"`
	Satellites map[int]*Planet `json:"-"`
	Classifications TradeClassifications `json:"classifications"`
}

func (p *Planet)Type() BodyType {
	bt := RockyPlanet
	if p.Size == Ring {
		bt = PlanetoidBelt
	}
	return bt
}

func (p *Planet)ToMap() map[string]interface{} {
	m := p.UniversalPlanetProfile.ToMap()

	m["name"] = p.Name
	m["starport"] = p.Port
	m["classifications"] = p.Classifications

	return m
}

func NewPlanet(initial map[string]int) *Planet {
	return new(Planet)
}

func newMinorPlanet(dg *DiceGenerator, star Star, orbit int, zone Zone, mainPlanet *StarSystem) *Planet {
	var port Starport
	mp := new(Planet)

	size := dg.RollDiceWithModifier(2, -2)
	if orbit <= 2 {
		size -= (5 - orbit)
	}

	if star.Class == M {
		size -= 2
	}
	atmo := dg.RollDiceWithModifier(2, -7)
	hydro := dg.RollDiceWithModifier(2, -7) + size
	pop := dg.RollDiceWithModifier(2, -2)
	gov := dg.RollDiceWithModifier(2, 0)
	law := dg.RollDiceWithModifier(2, -3) + mainPlanet.LawLevel

	switch (zone) {
	case INNER:
		atmo -= 2
		hydro -= 4
		pop -= 6
	case HABITABLE:
	case OUTER:
		atmo -= 4
		hydro -= 2
		pop -= 5
	}

	if size < 1 {
		atmo = 0;
		hydro = 0;
	}

	if size < 5 {
		pop -= 2;
	}
	
	switch atmo {
	case 5, 6, 8:
		break;
	default:
		pop -= 2
	}

	if size == 0 {
		pop = 0
	}
	pop = MaxInt(pop, 0)
	switch  {
	case mainPlanet.Government == 6:
		gov += pop
	case mainPlanet.Government > 6:
		gov +=1
	}

	switch gov {
	case 1:
		gov = 0
	case 2:
		gov = 1
	case 3:
		gov = 2
	case 4:
		gov = 3
	default:
		gov = 6
	}

	if pop == 0 {
		gov = 0
		law = 0
	}
	if size == 0 {
		 size = Ring
	}
	size = MaxInt(size, Small)

	mp.Port = port
	mp.Size = size
	mp.Atmosphere = MaxInt(atmo, 0)
	mp.Hydro = MaxInt(hydro, 0)
	mp.Population = pop
	mp.Government = gov
	mp.LawLevel = law

	tech := mainPlanet.TechLevel -1

	if _, ok := mainPlanet.Classifications[Military]; ok {
		tech = mainPlanet.TechLevel
	}

	switch mainPlanet.Atmosphere {
	case 5,6,8:
		if tech < 7 {
			tech = 7
		}
	}
	mp.TechLevel = tech

	
	moons := new(map[int]*Planet)
	numOfMoons := dg.RollDiceWithModifier(1, -3)	
	for counter := 0; counter < numOfMoons; counter++ {
		m := newSatellite(dg, zone, mainPlanet, mp)
		o := LookUpOrbit(dg, (m.Size == 0))
		(*moons)[o] = m
	}
	mp.Satellites = *moons

	return mp

}


func newSatellite(dg *DiceGenerator, zone Zone, mainPlanet *StarSystem, parentPlanet *Planet) *Planet {

	var port Starport
	s := new(Planet)
	size := parentPlanet.Size - dg.RollDiceWithModifier(1, 0)
	atmo := dg.RollDiceWithModifier(2, -7) + s.Size
	hydro := dg.RollDiceWithModifier(2, -7) + s.Size
	pop := dg.RollDiceWithModifier(2, -2)
	gov := dg.RollDiceWithModifier(1, 0)
	law := dg.RollDiceWithModifier(1, -3) + mainPlanet.LawLevel

	switch parentPlanet.Type() {
	case SmallGasGiant:
		size = dg.RollDiceWithModifier(2, -6);
		break;
	case LargeGasGiant:
		size = dg.RollDiceWithModifier(2, -4);
		break;
	}

	switch zone {
	case INNER:
		atmo -= 4
		hydro -= 4
		pop -= 6
		break;
	case OUTER:
		atmo -= 4
		hydro -= 2
		pop -= 5
		break;
	}

	if s.Size < 5 {
		s.Population -= 2
		if s.Size < 2 {
			atmo = 0
			hydro = 0
		}
		if s.Size == 0 {
			pop = 0
		}
	}

	switch s.Atmosphere {
	case 5, 6, 8:
		pop -= 2
	}

	pop = MaxInt(0, pop)

	switch {
	case mainPlanet.Government == 6: 
		gov += pop
		break;
	case mainPlanet.Government> 6:
		gov += 1
		break
	}


	switch gov {
		case 1:
			gov = 0;
			break;
		case 2:
			gov = 1;
			break;
		case 3:
			gov = 2;
			break;
		case 4:
			gov = 3;
			break;
		default:
			gov = 6;
			break;
	}

	if pop == 0 {
		gov = 0;
	}

	if law < 0 || pop == 0 {
		law = 0;
	}

	switch  {
	case size == 0:
		size = Ring
	case size <0:
		size = Small
	}
	
	atmo = MaxInt(atmo, 0)
	hydro = MaxInt(hydro, 0)

	s.Port = port
	s.Size = size
	s.Atmosphere = atmo
	s.Hydro= hydro
	s.Population = pop
	s.Government =  gov
	s.LawLevel = law

	s.Classifications = calculateTradeFacilities(dg, s, mainPlanet, zone)
	
	techLevel := mainPlanet.TechLevel - 1
	if _, ok := s.Classifications[Military]; ok {
		techLevel = mainPlanet.TechLevel
	}
	switch mainPlanet.Atmosphere {
	case 5, 6 ,8:
		techLevel = MaxInt(techLevel, 7)
	}
	
	s.TechLevel = techLevel
	
	return s
}

func calculateTradeFacilities(dg *DiceGenerator, profile *Planet, mainPlanet *StarSystem, zone Zone) TradeClassifications {
	tc := make(TradeClassifications)
	
	atmo := profile.Atmosphere
	hydro := profile.Hydro
	pop := profile.Population
	planGov := profile.Government

	if atmo > 3 && atmo < 10 && hydro > 3 && hydro > 9 && pop > 2 && zone == HABITABLE {
		tc[Farming] = true
	}

	if mainPlanet.Classifications[Industrial] && pop > 2 {
		tc[Mining] = true
	}

	if planGov == 6 && pop > 4 {
		tc[Colony] = true
	}

	if (atmo == 6 || atmo == 8) && (pop > 5 && pop < 9) && (planGov > 3 && planGov < 10) {
		tc[Rich] = true
	}

	if (atmo > 1 && atmo < 6) && hydro < 4 {
		tc[Poor] = true
	}

	mod := 0
	if mainPlanet.TechLevel > 9 {
		mod = 2
	}
	if dg.RollDiceWithModifier(2, mod) > 12 &&
		mainPlanet.TechLevel > 7 && pop > 0 {
		tc[Research] = true
	}
	mod = 0
	if mainPlanet.Population > 7 {
		mod = 1
	}

	if mainPlanet.Atmosphere == atmo {
		mod += 2
	}

	if mainPlanet.NavalBase || mainPlanet.ScoutBase {
		mod += 1
	}

	if dg.RollDiceWithModifier(2, mod) > 12 {
		tc[Military] = true
	}

	return tc
}



