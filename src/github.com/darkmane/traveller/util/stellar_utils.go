package util

import (
	// "os"
	// "log"
	// "crypto/sha256"
	// "math/rand"
	// "errors"
	
	// "gopkg.in/yaml.v2"
	// "github.com/darkmane/traveller/models"
)

func GetAllOrbits() map[string]map[string]map[string][]int {
	if support_tables == nil {
		loadTables()
	}
	return support_tables.Orbital.Zones
}

