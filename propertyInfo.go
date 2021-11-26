package main

type propertyInfo struct {
	address                                             string
	fairMarketValue                                     float64
	vacancyRate, managementRate, annualAppreciationRate float64 //MUST BE IN DECMIAL FORM! 5% WOULD BE 0.05
	numberOfUnits                                       int
	advertizingCostPerVacancy                           float64
}
