package main

type cashRequirementInfo struct {
	depositMadeWithOffer, lessProRationOfRents, cashRequiredToClose float64
}

func totalCashRequired(property cashRequirementInfo) float64 {
	return property.cashRequiredToClose + property.depositMadeWithOffer - property.lessProRationOfRents
}
