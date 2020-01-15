package main

import (
	"log"
	"net/http"
	"os"

	"github.com/darkmane/traveller/handlers"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"

	. "github.com/darkmane/traveller/util"
)

type Config struct {
	Database string `yaml:"database", envconfig:DATABASE`
	Username string `yaml:"username", envconfig:DBUSERNAME`
	Password string `yaml:"password", envconfig:DBPASSWORD`
	Seed     string `yaml:"seed", envconfig:RNG_SEED`
}

func main() {
	cfg := GetConfig()
	log.Printf("New Database: %v, Username: %v, Seed: %v", cfg.Database, cfg.Username, cfg.Seed)
	handlers.RegisterHandlers(http.HandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func GetConfig() Config {
	var cfg Config
	loadConfigFile(&cfg)
	readEnv(&cfg)
	return cfg
}
func loadConfigFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		ProcessError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		ProcessError(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("traveller", cfg)
	if err != nil {
		ProcessError(err)
	}
}
