package main


import "github.com/darkmane/traveller/handlers"
import (
	"log"
	"net/http"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Database string
	Username string
	Password string
}


func main() {
	handlers.RegisterHandlers(http.HandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func processError(err error){
	log.Fatal("Error: %s", err)
}

func loadConfig() Config {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}

	return cfg
}
