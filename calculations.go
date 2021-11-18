package main

import "fmt"

// Contains functions that do property calculations

func totalIncome(gorss_rents, parking, storage, laundry_vending, other float64) float64 {
	return gorss_rents + parking + storage + laundry_vending + other
}

func vacancyLoss(vacancy_rate, total_income float64) float64 {
	return (vacancy_rate / 100) * total_income
}

/*
func effectiveGrossIncome(total_income, vacancy_loss, float64) float64
{
	return total_income - vacancy_loss
}
*/

func main() {
	fmt.Println(vacancyLoss(5, 14400))
}
