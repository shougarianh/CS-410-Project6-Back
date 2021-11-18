package main

import (
	"fmt"
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
	//var percentage float64 = first_mtg_interest_rate / 100
	return 1 / 6
	//return first_mtg_total_principle * (math.Pow((1/6), (1+(percentage/2))) - 1) /// (1 - math.Pow((math.Pow((1+(percentage/2)), (1/6))), (first_mtg_amortization_period*(-12))))
}

func main() {
	fmt.Println(vacancyLoss(14400, 720))
	fmt.Println(firstMtgPrincipleBorrowed(105000, 25))
	fmt.Println(firstMtgTotalPrinciple(78750, 0))
	fmt.Println(firstMtgTotalMonthlyPayment(78750, 3.63, 30))
}
