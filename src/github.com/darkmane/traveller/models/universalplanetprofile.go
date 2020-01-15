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

func (upp *UniversalPlanetProfile) Load(init map[string]int) {
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
