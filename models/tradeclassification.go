package models

import (
	"bytes"
	"encoding/json"
)

type TradeClassification int

const (
	Agricultural TradeClassification = iota
	NonAgricultural
	Industrial
	NonIndustrial
	Rich
	Poor
	Farming
	Mining
	Colony
	Research
	Military
)

type TradeClassifications map[TradeClassification]bool

var tradeClassificationsToString = map[TradeClassification]string {
	Agricultural: "Agricultural",
	NonAgricultural: "NonAgricultural",
	Industrial: "Industrial",
	NonIndustrial: "NonIndustrial",
	Rich: "Rich",
	Poor: "Poor",
	Farming: "Farming",
	Mining: "Mining",
	Colony: "Colony",
	Research: "Research",
	Military: "Military",
}

var tradeClassificationsToID = map[string]TradeClassification {
	"Agricultural": Agricultural,
	"NonAgricultural": NonAgricultural,
	"Industrial": Industrial,
	"NonIndustrial": NonIndustrial,
	"Rich": Rich,
	"Poor": Poor,
	"Farming": Farming,
	"Mining": Mining,
	"Colony": Colony,
	"Research": Research,
	"Military": Military,
}

func (tc TradeClassification) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(tradeClassificationsToString[tc])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (tc *TradeClassification) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'B' in this case.
	*tc = tradeClassificationsToID[j]
	return nil
}

func (tc TradeClassifications) MarshalJSON() ([]byte, error) {
	ov := make([]TradeClassification, 0)
	for k, _ := range tc {
		ov = append(ov, k)
	}
	return json.Marshal(ov)
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (tc TradeClassifications) UnmarshalJSON(b []byte) error {
	var j []TradeClassification
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	for _, e := range j {
		tc[e] = true
	} 
	
	return nil
}