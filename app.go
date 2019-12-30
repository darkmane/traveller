package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/darkmane/traveller/handlers"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database string `yaml:"database", envconfig:DATABASE`
	Username string `yaml:"username", envconfig:DBUSERNAME`
	Password string `yaml:"password", envconfig:DBPASSWORD`
	Seed     string `yaml:"seed", envconfig:RNG_SEED`
}

func main() {
	cfg := GetConfig()
	fmt.Println("Database: %s, Username: %s, Seed: %s", cfg.Database, cfg.Username, cfg.Seed)
	handlers.RegisterHandlers(http.HandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func processError(err error) {
	log.Fatal("Error: %s", err)
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
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
