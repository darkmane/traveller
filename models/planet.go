package models

import (
	. "darkmane/traveller/util"
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

	name            string = "name"
	starport        string = "starport"
	classifications string = "classifications"
)

type Planet struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
	UniversalPlanetProfile
	Port      Starport `json:"starport"`
	ScoutBase bool     `json:"scout"`
	NavalBase bool     `json:"naval"`

	Satellites      map[int]*Planet      `json:"-"`
	Classifications TradeClassifications `json:"classifications"`

	stellarOrbit   int `json:"-"`
	planetaryOrbit int `json:"-"`
}

// Implement Orbit Interface

func (p *Planet) GetType() BodyType {
	bt := RockyPlanet
	if p.Size == Ring {
		bt = PlanetoidBelt
	}
	return bt
}

func (p *Planet) GetOrbit() (int, int) {
	return p.stellarOrbit, p.planetaryOrbit
}

// End Orbit interface

func (p *Planet) setOrbit(stellarOrbit *int, planetaryOrbit *int) {
	if stellarOrbit != nil {
		p.stellarOrbit = *stellarOrbit
	}

	if planetaryOrbit != nil {
		p.planetaryOrbit = *planetaryOrbit
	}
}

func (p *Planet) FromMap(init map[string]interface{}) {
	if p == nil {
		p = new(Planet)
	}
	for k, v := range init {
		switch k {
		case name:
			p.Name = v.(string)
			break
		case starport:
			s := []byte(v.(string))
			p.Port.UnmarshalJSON(s)
			break
		}
	}

	upp := p.UniversalPlanetProfile
	upp.FromMap(init)
	p.UniversalPlanetProfile = upp
}

func (p *Planet) ToMap() map[string]interface{} {
	if p == nil {
		return make(map[string]interface{})
	}

	m := p.UniversalPlanetProfile.ToMap()

	m[name] = p.Name
	m[starport] = p.Port

	return m
}

func NewPlanet(initial map[string]interface{}, dg *DiceGenerator) *Planet {
	p := new(Planet)

	// 	This checklist smrns generation of
	// the main world in a star system. 1. Determinelyltern prerena. 2. Checksystemconuno tat4e.
	// A. Findrtarpmtype. 8.Ch-k for naval bare. C.CheckforXWI bare. D.Check for gar giant.
	p.Port = generateStarport(dg)
	p.NavalBase = generateNavalBase(dg, p.Port)
	p.ScoutBase = generateScoutBase(dg, p.Port)
	p.Size = generateSize(dg)
	p.Atmosphere = generateAtmosphere(dg, p.Size)
	p.Hydro = generateHydrosphere(dg, p.Size, p.Atmosphere)
	p.Population = MaxInt(0, dg.RollDiceWithModifier(2, -2))
	p.Government = MaxInt(0, dg.RollDiceWithModifier(2, -7+p.Population))
	p.LawLevel = MaxInt(0, dg.RollDiceWithModifier(2, -7+p.Government))

	// 3. Name main world. 4.Decidsiftrawlzonemdsd. 5. Generate mainworld UPP.
	// A. Notertarpontyps.
	// 8. Main world rile: 2D-2.
	// C. Main world atmesphere: 2D-7 +
	// size. Ifsize0, thenatmorphere0.
	// D. Main world hydrographiu: ZD-7 +sire. If size I-,then hydrographiu 0;
	// if atmesphere 1- or A+. DM -4. If 1-1 than 0. then 0; i f greater than A, then A.
	// E. Populatian: 2D-2.
	// F. Government: 2D-7+populatim. G. Law Level: 2D.7+ povernrnent. H. Techndwical Iwel: 1D + DM9
	// from the tech level table.
	// 8. Notetra&elmrificationr.
	// 7. Recordrtatirticr for r e f e m . 8.Mapwrtem onrubrenormapgrid. 9. Establishcommunicationsroutes.
	return p
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

	switch zone {
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
		atmo = 0
		hydro = 0
	}

	if size < 5 {
		pop -= 2
	}

	switch atmo {
	case 5, 6, 8:
		break
	default:
		pop -= 2
	}

	if size == 0 {
		pop = 0
	}
	pop = MaxInt(pop, 0)
	switch {
	case mainPlanet.Government == 6:
		gov += pop
	case mainPlanet.Government > 6:
		gov += 1
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

	tech := mainPlanet.TechLevel - 1

	if _, ok := mainPlanet.Classifications[Military]; ok {
		tech = mainPlanet.TechLevel
	}

	switch mainPlanet.Atmosphere {
	case 5, 6, 8:
		if tech < 7 {
			tech = 7
		}
	}
	mp.TechLevel = tech

	moons := make(map[int]*Planet)
	numOfMoons := dg.RollDiceWithModifier(1, -3)
	for counter := 0; counter < numOfMoons; counter++ {
		m := newSatellite(dg, zone, mainPlanet, mp)
		o := LookUpOrbit(dg, (m.Size == 0))
		moons[o] = m
	}
	mp.Satellites = moons
	mp.setOrbit(&orbit, nil)
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

	switch parentPlanet.GetType() {
	case SmallGasGiant:
		size = dg.RollDiceWithModifier(2, -6)
		break
	case LargeGasGiant:
		size = dg.RollDiceWithModifier(2, -4)
		break
	}

	switch zone {
	case INNER:
		atmo -= 4
		hydro -= 4
		pop -= 6
		break
	case OUTER:
		atmo -= 4
		hydro -= 2
		pop -= 5
		break
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
		break
	case mainPlanet.Government > 6:
		gov += 1
		break
	}

	switch gov {
	case 1:
		gov = 0
		break
	case 2:
		gov = 1
		break
	case 3:
		gov = 2
		break
	case 4:
		gov = 3
		break
	default:
		gov = 6
		break
	}

	if pop == 0 {
		gov = 0
	}

	if law < 0 || pop == 0 {
		law = 0
	}

	switch {
	case size == 0:
		size = Ring
	case size < 0:
		size = Small
	}

	atmo = MaxInt(atmo, 0)
	hydro = MaxInt(hydro, 0)

	s.Port = port
	s.Size = size
	s.Atmosphere = atmo
	s.Hydro = hydro
	s.Population = pop
	s.Government = gov
	s.LawLevel = law

	s.Classifications = calculateTradeFacilities(dg, s, mainPlanet, zone)

	techLevel := mainPlanet.TechLevel - 1
	if _, ok := s.Classifications[Military]; ok {
		techLevel = mainPlanet.TechLevel
	}
	switch mainPlanet.Atmosphere {
	case 5, 6, 8:
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

func generateStarport(dg *DiceGenerator) Starport {
	var sp Starport
	switch dg.Roll() {
	case 2, 3, 4:
		sp = StarportA
	case 5, 6:
		sp = StarportB
	case 7, 8:
		sp = StarportC
	case 9:
		sp = StarportD
	case 10, 11:
		sp = StarportE
	case 12:
		sp = StarportX
	}
	return sp
}

func generateNavalBase(dg *DiceGenerator, sp Starport) bool {
	switch sp {
	case StarportA, StarportB:
		return (dg.Roll() > 7)
	default:
		return false
	}
}

func generateScoutBase(dg *DiceGenerator, sp Starport) bool {
	dm := 0
	switch sp {
	case StarportC:
		dm = -1
	case StarportB:
		dm = -2
	case StarportA:
		dm = -3
	case StarportE, StarportX:
		return false
	}

	return (dg.RollDiceWithModifier(2, dm) > 6)
}

func generateSize(dg *DiceGenerator) int {
	roll := dg.RollDiceWithModifier(2, -2)
	return MaxInt(roll, 0)
}

func generateAtmosphere(dg *DiceGenerator, size int) int {
	roll := dg.RollDiceWithModifier(2, -7+size)
	return MaxInt(0, roll)
}

func generateHydrosphere(dg *DiceGenerator, size int, atmos int) int {
	dm := 0
	switch {
	case atmos < 2:
		dm = -4
	case atmos > 9:
		dm = -4
	}
	if size < 2 {
		return 0
	} else {
		roll := dg.RollDiceWithModifier(2, dm)
		return MaxInt(MinInt(roll, 10), 0)
	}
}
