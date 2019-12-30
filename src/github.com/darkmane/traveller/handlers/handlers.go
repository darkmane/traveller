package handlers

import (
	"fmt"
	"net/http"
)

// Registers all Handlers for the Traveller REST API
//    -- StarSystem handlers
func RegisterHandlers(httpFunc func(pattern string, handler func(http.ResponseWriter, *http.Request))) {
	httpFunc("/", rootHandler)
	httpFunc("/starsystem", starSystemHandlers)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ROOT")
}
