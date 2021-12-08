package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Input struct {
	VacancyRate                 float64
	ManagementRate              float64
	AdvertizingCostPerVacancy   float64
	NumberOfUnits               int
	AnnualAppreciationRate      float64
	OfferPrice                  float64
	Repairs                     float64
	RepairsContigency           float64
	LenderFee                   float64
	BrokerFee                   float64
	Enviromentals               float64
	InspectionsOrEngineerReport float64
	Appraisals                  float64
	Misc                        float64
	TransferTax                 float64
	Legal                       float64
	DownPaymentPercent          float64
	FirstMtgInterestRate        float64
	FirstMtgAmortizationPeriod  float64
	FirstMtgCMHCFee             float64
	SecondMtgPrincipalAmount    float64
	SecondMtgInterestRate       float64
	SecondMtgAmortizationPeriod float64
	InterestOnlyPrincipleAmount float64
	OtherMonthlyFinancingCosts  float64
	InterestOnlyMonthlyPayment  float64
	GrossRents                  float64
	Parking                     float64
	Storage                     float64
	LaundryAndVending           float64
	Other                       float64
	PropertyTaxes               float64
	Insurance                   float64
	RepairsOperatingExpenses    float64
	Electricity                 float64
	Gas                         float64
	LawnAndSnowMaintence        float64
	WaterAndSewer               float64
	Cable                       float64
	Management                  float64
	Caretaking                  float64
	Advertizing                 float64
	AssociationFees             float64
	PestControl                 float64
	Security                    float64
	TrashRemoval                float64
	MiscOperatingExpenses       float64
	CommonAreaMaintence         float64
	CapitalImprovements         float64
	Accounting                  float64
	LegalOperatingExpenses      float64
	OtherOperatingExpenses      float64
	DepositMadeWithOffer        float64
	LessProRationOfRents        float64
}

type PropertyCalculations struct {
	TotalCashRequired       float64 `json:"TotalCashRequired"`
	AnnualProfitOrLoss      float64 `json:"AnnualProfitOrLoss"`
	CashFlowPerUnitPerMonth float64 `json:"CashFlowPerUnitPerMonth"`
	CashOnCashRoi           float64 `json:"CashOnCashRoi"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnPropertyInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnPropertyInfo")

	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var propInfo propertyInfo
	var purInfo purchaseInfo
	var finMonthInfo financingMonthlyInfo
	var inAnInfo incomeAnnualInfo
	var opExpenInfo operatingExpensesAnnualInfo
	var cashRecInfo cashRequirementInfo

	// property info ----
	propInfo.vacancyRate = input.VacancyRate
	propInfo.managementRate = input.ManagementRate
	propInfo.advertizingCostPerVacancy = input.AdvertizingCostPerVacancy
	propInfo.numberOfUnits = int(input.NumberOfUnits)
	propInfo.annualAppreciationRate = input.AnnualAppreciationRate

	// purchase info ----
	purInfo.offerPrice = input.OfferPrice
	purInfo.repairs = input.Repairs
	purInfo.repairsContigency = input.RepairsContigency
	purInfo.lenderFee = input.LenderFee
	purInfo.brokerFee = input.BrokerFee
	purInfo.enviromentals = input.Enviromentals
	purInfo.inspectionsOrEngineerReport = input.InspectionsOrEngineerReport
	purInfo.appraisals = input.Appraisals
	purInfo.misc = input.Misc
	purInfo.transferTax = input.TransferTax
	purInfo.legal = input.Legal

	real_purchase_price := realPurchasePrice(purInfo)

	// financing monthly
	finMonthInfo.downPaymentPercent = input.DownPaymentPercent
	finMonthInfo.firstMtgInterestRate = input.FirstMtgInterestRate
	finMonthInfo.firstMtgAmortizationPeriod = input.FirstMtgAmortizationPeriod
	finMonthInfo.firstMtgCMHCFee = input.FirstMtgCMHCFee
	second_mtg_principal_amount := input.SecondMtgPrincipalAmount
	second_mtg_interest_rate := input.SecondMtgInterestRate
	second_mtg_amortization_period := input.SecondMtgAmortizationPeriod
	interest_only_principle_amount := input.InterestOnlyPrincipleAmount
	first_mtg_principle_borrowed := firstMtgPrincipleBorrowed(purInfo.offerPrice, finMonthInfo.downPaymentPercent)
	first_mtg_total_principle := firstMtgTotalPrinciple(first_mtg_principle_borrowed, finMonthInfo.firstMtgCMHCFee)
	first_mtg_total_monthly_payment := firstMtgTotalMonthlyPayment(first_mtg_total_principle, finMonthInfo.firstMtgInterestRate, finMonthInfo.firstMtgAmortizationPeriod)
	second_mtg_total_monthly_payment := secondMtgTotalMonthlyPayment(second_mtg_principal_amount, second_mtg_interest_rate, second_mtg_amortization_period)
	cash_required_to_close := cashRequiredToClose(real_purchase_price, first_mtg_principle_borrowed, second_mtg_principal_amount, interest_only_principle_amount)
	other_monthly_financing_costs := input.OtherMonthlyFinancingCosts
	interest_only_monthly_payment := input.InterestOnlyMonthlyPayment

	// income annual

	inAnInfo.grossRents = input.GrossRents
	inAnInfo.parking = input.Parking
	inAnInfo.storage = input.Storage
	inAnInfo.laundryAndVending = input.LaundryAndVending
	inAnInfo.other = input.Other
	total_income := totalIncome(inAnInfo)
	vacancy_loss := vacancyLoss(propInfo.vacancyRate, total_income)
	effective_gross_income := effectiveGrossIncome(total_income, vacancy_loss)

	// Operating expenses

	opExpenInfo.propertyTaxes = input.PropertyTaxes
	opExpenInfo.insurance = input.Insurance
	opExpenInfo.repairs = input.RepairsOperatingExpenses
	opExpenInfo.electricity = input.Electricity
	opExpenInfo.gas = input.Gas
	opExpenInfo.lawnAndSnowMaintence = input.LawnAndSnowMaintence
	opExpenInfo.waterAndSewer = input.WaterAndSewer
	opExpenInfo.cable = input.Cable
	opExpenInfo.manangement = input.Management
	opExpenInfo.caretaking = input.Caretaking
	opExpenInfo.advertizing = input.Advertizing
	opExpenInfo.associationFees = input.AssociationFees
	opExpenInfo.pestControl = input.PestControl
	opExpenInfo.security = input.Security
	opExpenInfo.trashRemoval = input.TrashRemoval
	opExpenInfo.misc = input.MiscOperatingExpenses
	opExpenInfo.commonAreaMaintence = input.CommonAreaMaintence
	opExpenInfo.capitalImprovements = input.CapitalImprovements
	opExpenInfo.accounting = input.Accounting
	opExpenInfo.legal = input.LegalOperatingExpenses
	opExpenInfo.other = input.OtherOperatingExpenses
	evictions := evictions(propInfo.numberOfUnits, propInfo.vacancyRate)

	total_expenses := totalExpenses(opExpenInfo) + evictions

	// net operating income

	net_operating_income := netOperatingIncome(effective_gross_income, total_expenses)

	// cash requirments

	cashRecInfo.depositMadeWithOffer = input.DepositMadeWithOffer
	cashRecInfo.lessProRationOfRents = input.LessProRationOfRents
	cashRecInfo.cashRequiredToClose = cash_required_to_close - cashRecInfo.depositMadeWithOffer
	total_cash_required := totalCashRequired(cashRecInfo)

	// cash flow summary

	debt_servicing_cost := debtServiceCosts(first_mtg_total_monthly_payment, second_mtg_total_monthly_payment, interest_only_monthly_payment, other_monthly_financing_costs)
	annual_profit_or_loss := annualProfitOrLoss(net_operating_income, debt_servicing_cost)
	total_monthly_profit_loss := totalMonthlyProfitLoss(annual_profit_or_loss)
	cash_flow_per_unit_per_month := cashFlowPerUnitPerMonth(total_monthly_profit_loss, float64(propInfo.numberOfUnits))

	// quick analysis
	cash_on_cash_roi := cashOnCashRoi(total_cash_required, annual_profit_or_loss)
	propCalcs := PropertyCalculations{TotalCashRequired: total_cash_required, AnnualProfitOrLoss: annual_profit_or_loss,
		CashOnCashRoi: cash_on_cash_roi, CashFlowPerUnitPerMonth: cash_flow_per_unit_per_month}
	json.NewEncoder(w).Encode(propCalcs)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/propertyInfo", returnPropertyInfo)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
