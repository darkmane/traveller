package main

import (
	"log"
	"net/http"
	"github.com/darkmane/traveller/handlers"

	. "github.com/darkmane/traveller/util"
)

func main() {
	cfg := GetConfig()
	log.Printf("New Database: %v, Username: %v, Seed: %v", cfg.Database, cfg.Username, cfg.Seed)
	handlers.RegisterHandlers(http.HandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}



