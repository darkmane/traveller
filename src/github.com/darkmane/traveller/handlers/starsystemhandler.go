package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/darkmane/traveller/models"
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
	init, err := parseRequest(r)
	log.Printf(fmt.Sprintf("Body: %v", init))
	if err != nil {
		w.WriteHeader(500)
		log.Printf(fmt.Sprintf("%v", err))
		return
	}
	ss := new(models.StarSystem)
	ss.Load(init)
	results, err := json.Marshal(ss)
	if err != nil {
		w.WriteHeader(500)
		log.Printf(fmt.Sprintf("%v", err))
		return
	}
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(results)
}

func getStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	ss := new(models.StarSystem)
	results, err := json.Marshal(ss)
	if err != nil {
		w.WriteHeader(500)
		log.Printf(fmt.Sprintf("%v", err))
		return
	}
	log.Printf(fmt.Sprintf("StarSystem: %v", ss))
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(results)
}

func getMultipleStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	upp := new(models.UniversalPlanetProfile)
	var upps []models.UniversalPlanetProfile
	upps = make([]models.UniversalPlanetProfile, 1)
	upps[0] = *upp
	results, err := json.Marshal(upps)
	if err != nil {
		w.WriteHeader(500)
		log.Printf(fmt.Sprintf("%v", err))
		return
	}
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(results)
}

func updateStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	upp := new(models.UniversalPlanetProfile)
	results, err := json.Marshal(upp)
	if err != nil {
		w.WriteHeader(500)
		log.Printf(fmt.Sprintf("%v", err))
		return
	}
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.Write(results)
}

func deleteStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("deleteStarSystemHandler")
}
