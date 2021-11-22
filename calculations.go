package main

import (
	"fmt"
	"math"
)

// Contains functions that do property calculations

func totalIncome(gorss_rents, parking, storage, laundry_vending, other float64) float64 {
	return gorss_rents + parking + storage + laundry_vending + other
}

func vacancyLoss(vacancy_rate, total_income float64) float64 {
	return (vacancy_rate / 100) * total_income
}

func effectiveGrossIncome(total_income, vacancy_loss float64) float64 {
	return total_income - vacancy_loss
}

func firstMtgPrincipleBorrowed(offer_price, down_payment_percent float64) float64 {
	return (offer_price - (offer_price*down_payment_percent)/100)

}

func firstMtgTotalPrinciple(first_mtg_principle_borrowed, first_mtg_CMHC_fee float64) float64 {
	return first_mtg_principle_borrowed * (1 + first_mtg_CMHC_fee)
}

func firstMtgTotalMonthlyPayment(first_mtg_total_principle, first_mtg_interest_rate, first_mtg_amortization_period float64) float64 {
	var percentage float64 = first_mtg_interest_rate / 100.0
	return first_mtg_total_principle * (math.Pow((1+(percentage/2)), (1.0/6.0)) - 1) / (1 - math.Pow((math.Pow((1+(percentage/2)), (1.0/6.0))), (first_mtg_amortization_period*(-12))))
}

func secondMtgTotalMonthlyPayment(second_mtg_principal_amount, second_mtg_interest_rate, second_mtg_amortization_period float64) float64 {
	var percentage float64 = second_mtg_interest_rate / 100
	return second_mtg_principal_amount * (math.Pow((1+(percentage/2)), (1.0/6.0)) - 1) / (1 - math.Pow((math.Pow((1+(percentage/2)), (1.0/6.0))), (second_mtg_amortization_period*(-12))))
}

func interestTotalMonthlyPayment(interest_only_principle_amount, interest_only_interest_rate float64) float64 {
	return interest_only_principle_amount * (interest_only_interest_rate / 100) / 12
}

func debtServiceCosts(first_mtg_total_monthly_payment, second_mtg_total_monthly_payment, interest_only_monthly_payment, other_monthly_financing_costs float64) float64 {
	return (-first_mtg_total_monthly_payment - second_mtg_total_monthly_payment - interest_only_monthly_payment - other_monthly_financing_costs) * 12
}

func annualProfitOrLoss(net_operating_income, debt_servicing_cost float64) float64 {
	return net_operating_income + debt_servicing_cost
}

func totalMonthlyProfitLoss(anual_profit_or_loss float64) float64 {
	return anual_profit_or_loss / 12
}

func cashFlowPerUnitPerMonth(total_monthly_profit_loss, number_of_units float64) float64 {
	return total_monthly_profit_loss / number_of_units
}

func main() {
	fmt.Println(firstMtgTotalMonthlyPayment(78750, 3.63, 30))
	fmt.Println(debtServiceCosts(357.94, 0, 0, 0))
	fmt.Println(totalMonthlyProfitLoss(4604))

}
