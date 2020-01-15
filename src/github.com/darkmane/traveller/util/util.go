package util

import (
	"os"
	"log"
	"crypto/sha256"
	"math/rand"
	"errors"

	"gopkg.in/yaml.v2"
)

const sides int = 6

var support_tables *tables

type DiceGenerator struct {
	rng *rand.Rand
}

// Create a utility class for rolling dice
func NewDiceGenerator(seed string) DiceGenerator {
	h := sha256.New()
	h.Write([]byte(seed))

	bs := h.Sum(nil)
	var v int64 = 0
	for _, element := range bs {
		v += int64(element)
	}
	s := rand.NewSource(v)
	return DiceGenerator{rng: rand.New(s)}

}

// Roll an arbitrary number of d6 with a modifer at the end.
func (dg *DiceGenerator) RollDiceWithModifier(dice int, mod int) int {
	v := 0
	for index := 1; index <= dice; index++ {
		v = dg.rng.Intn(sides)
	}
	return v + mod
}

// Roll a specified number of dice
func (dg *DiceGenerator) RollDice(dice int) int {
	return dg.RollDiceWithModifier(dice, 0)
}

// Roll 2d6
func (dg *DiceGenerator) Roll() int {
	return dg.RollDice(2)
}


type tables struct {
	Conversions map[string]map[string]float64 `json:"conversions"`
	Orbital orbitalTable `json:"orbital"`
	Luminosity map[string]map[string]float32 `json:"luminosity"` 
	Profile profileTable `json:"profile"`
}

type orbitalTable struct {
	Zones map[string]map[string]map[string][]int `json:"zones"`
	Distances []float32 `json:"distances"`
	SatelliteOrbits map[string]map[int]int `json:"satellite"`
}

type profileTable struct {
	Indices map[string][]string `json:"indices"`
	Formatting map[string]map[string]string `json:"formatting"`
}

func loadTables() {
	f, err := os.Open("tables.yml")
	if err != nil {
		ProcessError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&support_tables)
	if err != nil {
		ProcessError(err)
	}
}

func GetSatelliteOrbits() map[string]map[int]int {
	if support_tables == nil {
		loadTables()
	}
	return support_tables.Orbital.SatelliteOrbits
}

func LookUpOrbit(dg *DiceGenerator, ring bool) int{
	satellite_orbits := GetSatelliteOrbits()
	key := "extreme"
	var roll int
	roll = dg.Roll()
	switch {
	case roll <= 7:
		key = "close"
	case roll <= 11:
		key = "far"
	}
	var rv int
	if options, ok := satellite_orbits[key]; ok {
		roll = dg.Roll()
		if len(options) == 6 {
			roll = dg.RollDiceWithModifier(1,0)
		}
		rv =  options[roll]
	} else {
		ProcessError(errors.New("Unknown Key for Satellite Orbits"))
	}
	return rv
}

func ProcessError(err error) {
	log.Fatal("Error: %s", err)
}

func MaxInt(nums ...int) int{
	max := nums[0]
	for _,num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}