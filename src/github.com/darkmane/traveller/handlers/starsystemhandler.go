package handlers

import (
	"fmt"
	"net/http"
)

func starSystemHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getStarSystemHandler(w, r)
	case "POST":
		createStarSystemHandler(w, r)
	case "PUT":
		updateStarSystemHandler(w, r)
	case "DELETE":
		deleteStarSystemHandler(w, r)
	}

}

func createStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "createStarSystemHandler")
}

func getStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "getStarSystemHandler")
}

func getMultipleStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "getMultipleStarSystemHandler")
}

func updateStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "updateStarSystemHandler")
}

func deleteStarSystemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "deleteStarSystemHandler")
}
