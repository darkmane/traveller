package util

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// Config Basic application configuration
type Config struct {
	LogLevel string `yaml:"loglevel" envconfig:"LOG_LEVEL"`
	Database string `yaml:"database" envconfig:"DATABASE"`
	Username string `yaml:"username" envconfig:"DBUSERNAME"`
	Password string `yaml:"password" envconfig:"DBPASSWORD"`
	Seed     string `yaml:"seed" envconfig:"RNG_SEED"`
}

// GetConfig Load App configuration
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
