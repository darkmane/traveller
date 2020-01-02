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
