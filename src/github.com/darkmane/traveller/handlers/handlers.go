package handlers

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"

	. "github.com/darkmane/traveller/util"
)
var cfg Config


const (
	GET string = "GET"
	POST string = "POST"
	DELETE string = "DELETE"
	PUT string = "PUT"
	CONTENT_TYPE string = "Content-Type"
	APPLICATION_JSON string = "application/json"
)

// Registers all Handlers for the Traveller REST API
//    -- StarSystem handlers
func RegisterHandlers(httpFunc func(pattern string, handler func(http.ResponseWriter, *http.Request))) {
	cfg = GetConfig()
	httpFunc("/starsystem", starSystemHandlers)
	httpFunc("/", rootHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ROOT")
}

func parseRequest(r *http.Request) (map[string]interface{}, error) {
	var rv map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rv)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %v", err))
		return nil, err
	}

	return rv, nil
}
