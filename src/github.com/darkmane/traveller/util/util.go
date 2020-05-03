package util

import (
	"log"
	// "fmt"
)

const sides int = 6

var supportTables *tables

type tables struct {
	Conversions map[string]map[string]float64 `json:"conversions"`
	Orbital     orbitalTable                  `json:"orbital"`
	Luminosity  map[string]map[string]float32 `json:"luminosity"`
	Profile     profileTable                  `json:"profile"`
}

type orbitalTable struct {
	Zones           map[string]map[string]map[string][]int `json:"zones"`
	Distances       []float32                              `json:"distances"`
	SatelliteOrbits map[string]map[string]int              `json:"satellite"`
}

type profileTable struct {
	Indices    map[string][]string          `json:"indices"`
	Formatting map[string]map[string]string `json:"formatting"`
}

// ProcessError Fatal errors
func ProcessError(err error) {
	log.Fatal("Error: ", err)
}

// MaxInt get the maximum integer value
func MaxInt(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

// MinInt get the minimum integer value
func MinInt(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	return min
}

// Interface2Int Convert interface{} to integer
func Interface2Int(intf interface{}) int {
	t := int(intf.(float64))
	return t
}

// Intersect Get the intersection of 2 maps
func Intersect(a map[int]interface{}, b map[int]interface{}) map[int]interface{} {
	intersection := make(map[int]interface{})

	for k, v := range a {
		if _, ok := b[k]; ok {
			intersection[k] = v
		}
	}
	return intersection
}
