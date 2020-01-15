package models

type UniversalPlanetProfile struct {
	Size       int `json:"size"`
	Atmosphere int `json:"atmosphere"`
	Hydro      int `json:"hydro"`
	Population int `json:"population"`
	Government int `json:"government"`
	LawLevel   int `json:"lawlevel"`
	TechLevel  int `json:"techlevel"`
}

func (upp *UniversalPlanetProfile) FromMap(init map[string]int) {
	if val, ok := init["size"]; ok {
		upp.Size = val
	}
	if val, ok := init["atmosphere"]; ok {
		upp.Atmosphere = val
	}
	if val, ok := init["hydro"]; ok {
		upp.Hydro = val
	}
	if val, ok := init["population"]; ok {
		upp.Population = val
	}
	if val, ok := init["government"]; ok {
		upp.Government = val
	}
	if val, ok := init["lawlevel"]; ok {
		upp.LawLevel = val
	}
	if val, ok := init["techlevel"]; ok {
		upp.TechLevel = val
	}
}

func (upp *UniversalPlanetProfile) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["size"] = upp.Size
	m["atmosphere"] = upp.Atmosphere
	m["hydro"] = upp.Hydro
	m["population"] = upp.Population
	m["government"] = upp.Government
	m["lawlevel"] = upp.LawLevel
	m["techlevel"] = upp.TechLevel

	return m
}