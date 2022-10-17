package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"darkmane/traveller/models"
	. "darkmane/traveller/util"
)

func starSystemHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case GET:
		getStarSystemHandler(w, r)
	case POST:
		createStarSystemHandler(w, r)
	case PUT:
		updateStarSystemHandler(w, r)
	case DELETE:
		deleteStarSystemHandler(w, r)
	}

}

func createStarSystemHandler(w http.ResponseWriter, r *http.Request) {

	// t := time.Now()
	// dg := NewDiceGenerator(fmt.Sprintf("%d", t.Unix()))
	init, err := parseRequest(r)
	log.Trace().Interface("initial_map", init)
	if err != nil {
		w.WriteHeader(500)
		log.Error().Err(err)
		return
	}
	ss := new(models.StarSystem)
	ss.FromMap(init)
	results, err := json.Marshal(ss)
	if err != nil {
		w.WriteHeader(500)
		log.Error().Err(err)
		return
	}
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(results)
}

func getStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	// ss := new(models.StarSystem)
	// dg := NewDiceGenerator(cfg.Seed)
	t := time.Now()
	dg := NewDiceGenerator(fmt.Sprintf("%d", t.Unix()))

	ss := models.NewStarSystem(make(map[string]interface{}), &dg)
	if ss.Planet == nil {
		log.Debug().Interface("star_system", ss).Msg("Planet is nil")
	}

	results, err := json.Marshal(ss)
	if err != nil {
		w.WriteHeader(500)
		log.Error().Err(err).Msg("Unable to Marshal StarSystem")
		return
	}
	log.Trace().Interface("star_system", ss)
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(results)
}

func getMultipleStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	// dg := NewDiceGenerator(cfg.Seed)
	upp := new(models.UniversalPlanetProfile)
	var upps []models.UniversalPlanetProfile
	upps = make([]models.UniversalPlanetProfile, 1)
	upps[0] = *upp
	results, err := json.Marshal(upps)
	if err != nil {
		w.WriteHeader(500)
		log.Error().Err(err)

		return
	}
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(results)
}

func updateStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	// dg := NewDiceGenerator(cfg.Seed)
	upp := new(models.UniversalPlanetProfile)
	results, err := json.Marshal(upp)
	if err != nil {
		w.WriteHeader(500)
		log.Error().Err(err)
		return
	}
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(results)
}

func deleteStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	// dg := NewDiceGenerator(cfg.Seed)
	log.Trace().Msg("deleteStarSystemHandler")
}
