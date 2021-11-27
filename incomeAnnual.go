package main

type incomeAnnualInfo struct {
	grossRents, parking, storage, laundryAndVending, other float64
}

func totalIncome(property incomeAnnualInfo) float64 {
	return property.grossRents + property.parking + property.storage + property.laundryAndVending + property.other
}

func vacancyLoss(vacancy_rate, total_income float64) float64 {
	return (vacancy_rate / 100) * total_income
}

func effectiveGrossIncome(total_income, vacancy_loss float64) float64 {
	return total_income - vacancy_loss
}
