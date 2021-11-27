package main

type operatingExpensesAnnualInfo struct {
	propertyTaxes, insurance, repairs, electricity, gas, lawnAndSnowMaintence           float64
	waterAndSewer, cable, manangement, caretaking, advertizing, associationFees         float64
	pestControl, security, trashRemoval, misc, commonAreaMaintence, capitalImprovements float64
	accounting, legal, badDebts, other                                                  float64
}

func evictions(number_of_units int, vacancy_rate float64) float64 {
	return float64(number_of_units) * 12 * (vacancy_rate / 100) / 2 / 10 * 1000
}

func totalExpenses(property operatingExpensesAnnualInfo) float64 {
	var sum float64 = 0
	sum += property.propertyTaxes + property.insurance + property.repairs + property.electricity + property.gas + property.lawnAndSnowMaintence
	sum += property.waterAndSewer + property.cable + property.manangement + property.caretaking + property.advertizing + property.associationFees
	sum += property.pestControl + property.security + property.trashRemoval + property.misc + property.commonAreaMaintence + property.capitalImprovements
	sum += property.accounting + property.legal + property.badDebts + property.other
	return sum
}
