package main

func main() {
	var propInfo propertyInfo
	var purInfo purchaseInfo
	var finMonthInfo financingMonthlyInfo
	var inAnInfo incomeAnnualInfo
	var opExpenInfo operatingExpensesAnnualInfo


	// property info ----
	propInfo.vacancyRate = 5
	propInfo.managementRate = 5
	propInfo.advertizingCostPerVacancy = 100
	propInfo.numberOfUnits = 1
	propInfo.annualAppreciationRate = 3

	// purchase info ----
	purInfo.offerPrice = 105000
	purInfo.repairs = 0
	purInfo.repairsContigency = 0
	purInfo.lenderFee = 1210
	purInfo.brokerFee = 0
	purInfo.enviromentals = 0
	purInfo.inspectionsOrEngineerReport = 0
	purInfo.appraisals = 675
	purInfo.misc = 1567
	purInfo.transferTax = 1227.36
	purInfo.legal = 950

	real_purchase_price = realPurchasePrice(purInfo)

	// financing monthly
	finMonthInfo.downPaymentPercent = 25;
	finMonthInfo.firstMtgInterestRate = 3.63;
	finMonthInfo.firstMtgAmortizationPeriod = 30;
	finMonthInfo.firstMtgCMHCFee = 0;
	second_mtg_principal_amount = 0
	interest_only_principle_amount = 0
	first_mtg_principle_borrowed = firstMtgPrincipleBorrowed(purInfo.offerPrice, finMonthInfo.downPaymentPercent)
	first_mtg_total_principle = firstMtgTotalPrinciple(first_mtg_principle_borrowed, finMonthInfo.firstMtgCMHCFee)
	first_mtg_total_monthlyPayment = firstMtgTotalMonthlyPayment(first_mtg_total_principle, finMonthInfo.firstMtgInterestRate, finMonthInfo.firstMtgAmortizationPeriod)
	cash_required_to_close = cashRequiredToClose(real_purchase_price, first_mtg_principle_borrowed, second_mtg_principal_amount, interest_only_principle_amount)

	// income annual

	inAnInfo.grossRents = 14400
	inAnInfo.parking = 0
	inAnInfo.storage = 0
	inAnInfo.laundryAndVending = 0
	inAnInfo.other = 0
	total_income = totalIncome(inAnInfo)
	vacancy_loss = vacancyLoss(propInfo.vacancyRate ,total_income)
	effective_gross_income = effectiveGrossIncome(total_income, vacancy_loss)

	// Operating expenses

	opExpenInfo.propertyTaxes = 600
	opExpenInfo.insurance = 504
	opExpenInfo.repairs = 720
	opExpenInfo.electricity = 0
	opExpenInfo.gas = 0
	opExpenInfo.lawnAndSnowMaintence = 0
	opExpenInfo.waterAndSewer = 0
	opExpenInfo.cable = 0
	opExpenInfo.manangement = 720
	opExpenInfo.caretaking = 0
	opExpenInfo.advertizing = 30
	opExpenInfo.associationFees = 2016
	opExpenInfo.pestControl = 140
	opExpenInfo.security = 20
	opExpenInfo.trashRemoval = 0
	opExpenInfo.misc = 0
	opExpenInfo.commonAreaMaintence = 0
	opExpenInfo.capitalImprovements = 0
	opExpenInfo.accounting = 0
	opExpenInfo.legal = 0
	opExpenInfo.other = 0
	opExpenInfo.evictions = 30

	total_expenses = total_expenses(opExpenInfo)


	



}

func realPurchasePrice(purInfo invalid type) {
	panic("unimplemented")
}