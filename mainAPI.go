package main

import "fmt"

func main() {
	var propInfo propertyInfo
	var purInfo purchaseInfo
	var finMonthInfo financingMonthlyInfo
	var inAnInfo incomeAnnualInfo
	var opExpenInfo operatingExpensesAnnualInfo
	var cashRecInfo cashRequirementInfo

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

	real_purchase_price := realPurchasePrice(purInfo)

	// financing monthly
	finMonthInfo.downPaymentPercent = 25.0
	finMonthInfo.firstMtgInterestRate = 3.63
	finMonthInfo.firstMtgAmortizationPeriod = 30
	finMonthInfo.firstMtgCMHCFee = 0.0
	second_mtg_principal_amount := 0.0
	second_mtg_interest_rate := 12.0
	second_mtg_amortization_period := 9999.99
	interest_only_principle_amount := 0.0
	first_mtg_principle_borrowed := firstMtgPrincipleBorrowed(purInfo.offerPrice, finMonthInfo.downPaymentPercent)
	first_mtg_total_principle := firstMtgTotalPrinciple(first_mtg_principle_borrowed, finMonthInfo.firstMtgCMHCFee)
	first_mtg_total_monthly_payment := firstMtgTotalMonthlyPayment(first_mtg_total_principle, finMonthInfo.firstMtgInterestRate, finMonthInfo.firstMtgAmortizationPeriod)
	second_mtg_total_monthly_payment := secondMtgTotalMonthlyPayment(second_mtg_principal_amount, second_mtg_interest_rate, second_mtg_amortization_period)
	cash_required_to_close := cashRequiredToClose(real_purchase_price, first_mtg_principle_borrowed, second_mtg_principal_amount, interest_only_principle_amount)
	other_monthly_financing_costs := 0.0
	interest_only_monthly_payment := 0.0

	// income annual

	inAnInfo.grossRents = 14400.0
	inAnInfo.parking = 0.0
	inAnInfo.storage = 0.0
	inAnInfo.laundryAndVending = 0.0
	inAnInfo.other = 0.0
	total_income := totalIncome(inAnInfo)
	vacancy_loss := vacancyLoss(propInfo.vacancyRate, total_income)
	effective_gross_income := effectiveGrossIncome(total_income, vacancy_loss)

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
	evictions := evictions(propInfo.numberOfUnits, propInfo.vacancyRate)

	total_expenses := totalExpenses(opExpenInfo) + evictions

	// net operating income

	net_operating_income := netOperatingIncome(effective_gross_income, total_expenses)

	// cash requirments

	cashRecInfo.depositMadeWithOffer = 0
	cashRecInfo.lessProRationOfRents = 0
	cashRecInfo.cashRequiredToClose = cash_required_to_close - cashRecInfo.depositMadeWithOffer
	total_cash_required := totalCashRequired(cashRecInfo)

	// cash flow summary

	debt_servicing_cost := debtServiceCosts(first_mtg_total_monthly_payment, second_mtg_total_monthly_payment, interest_only_monthly_payment, other_monthly_financing_costs)
	annual_profit_or_loss := annualProfitOrLoss(net_operating_income, debt_servicing_cost)
	total_monthly_profit_loss := totalMonthlyProfitLoss(annual_profit_or_loss)
	cash_flow_per_unit_per_month := cashFlowPerUnitPerMonth(total_monthly_profit_loss, float64(propInfo.numberOfUnits))

	// quick analysis
	cash_on_cash_roi := cashOnCashRoi(total_cash_required, annual_profit_or_loss)
	fmt.Println(total_cash_required)
	fmt.Println(annual_profit_or_loss)
	fmt.Println(cash_flow_per_unit_per_month)
	fmt.Println(cash_on_cash_roi)
}
