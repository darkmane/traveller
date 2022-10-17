package models

import (
	"fmt"

	"github.com/darkmane/traveller/util"
	"github.com/rs/zerolog/log"
)

type UniversalPlanetProfile struct {
	Size       int `json:"size"`
	Atmosphere int `json:"atmosphere"`
	Hydro      int `json:"hydro"`
	Population int `json:"population"`
	Government int `json:"government"`
	LawLevel   int `json:"lawlevel"`
	TechLevel  int `json:"techlevel"`
}

func (upp *UniversalPlanetProfile) FromMap(init map[string]interface{}) {
	for k, v := range init {
		switch k {
		case "size":
			upp.Size = util.Interface2Int(v)
			break
		case "atmosphere":
			upp.Atmosphere = util.Interface2Int(v)
			break
		case "hydro":
			upp.Hydro = util.Interface2Int(v)
			break
		case "population":
			upp.Population = util.Interface2Int(v)
			break
		case "government":
			upp.Government = util.Interface2Int(v)
			break
		case "lawlevel":
			upp.LawLevel = util.Interface2Int(v)
			break
		case "techlevel":
			upp.TechLevel = util.Interface2Int(v)
			break
		default:
			log.Debug().Msg(fmt.Sprintf("%s: %v", k, v))
			break
		}
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
