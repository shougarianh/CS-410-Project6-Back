package main

type purchaseInfo struct {
	offerPrice, repairs, repairsContigency, lenderFee, brokerFee                     float64 //All values are floats
	enviromentals, inspectionsOrEngineerReport, appraisals, misc, transferTax, legal float64
}

func realPurchasePrice(property purchaseInfo) float64 {
	var sum float64 = 0
	sum += property.offerPrice + property.repairs + property.repairsContigency
	sum += property.lenderFee + property.brokerFee + property.enviromentals
	sum += property.inspectionsOrEngineerReport + property.appraisals + property.misc
	sum += property.transferTax + property.legal
	return sum
}
