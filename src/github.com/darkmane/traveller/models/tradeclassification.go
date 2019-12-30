package models

type TradeClassifications int {
	Agricultural = iota
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
}

type TradeClassification struct {
	Body int64
	Classification TradeClassifications
}