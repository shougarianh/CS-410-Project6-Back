package main

type cashFlowSummaryInfo struct {
	//we can get these values from the other files.
	//effectiveGrossIncome from incomeAnnual.go, operatingExpenses from operatingExpenses.go
	//netOperatingIncome from netOperatingIncome.go
	effectiveGrossIncome, incomeAnnual, netOperatingIncome float64
}

func debtServiceCosts(first_mtg_total_monthly_payment, second_mtg_total_monthly_payment, interest_only_monthly_payment, other_monthly_financing_costs float64) float64 {
	return (-first_mtg_total_monthly_payment - second_mtg_total_monthly_payment - interest_only_monthly_payment - other_monthly_financing_costs) * 12
}

func annualProfitOrLoss(net_operating_income, debt_servicing_cost float64) float64 {
	return net_operating_income + debt_servicing_cost
} //questioning how to go about this
// use cashFlowSummaryInfo struct and the function I made above?

func totalMonthlyProfitLoss(anual_profit_or_loss float64) float64 {
	return anual_profit_or_loss / 12
}

func cashFlowPerUnitPerMonth(total_monthly_profit_loss, number_of_units float64) float64 {
	return total_monthly_profit_loss / number_of_units
}
