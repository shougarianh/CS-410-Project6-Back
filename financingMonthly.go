package main

import (
	"math"
)

type financingMonthlyInfo struct {
	firstMtgInterestRate, downPaymentPercent, firstMtgCMHCFee float64 //Percentages; so 3.63% would be 0.0363
	firstMtgAmortizationPeriod                                int
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

func cashRequiredToClose(real_purchase_price, first_mtg_principle_borrowed, second_mtg_principal_amount, interest_only_principle_amount float64) float64 {
	return real_purchase_price - first_mtg_principle_borrowed - second_mtg_principal_amount - interest_only_principle_amount
}

//some formulas were not included such as 2ndMtgPrincipleAmounts because they were not given. Lots of missing info for other values here
//do not know if those values matter as much
