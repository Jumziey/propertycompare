package propertycost

import (
	"errors"
	"math"
)

//### NOTES ####
//PropertyTax (actual value)
//Calculated as
//PropertyTaxValue = 75 percent of price
//(not a true value, normally its decided on the average costs
// of properties in your neighbourhood over the last 3 years.)
//PropertyTax = min(8349, PropertyTaxValue*0.0075)
//i.e. everything above 1 484 267(sek)
//

//DownPayment describes a down payment and
//its assocciated costs.
//AmountInHand - how much money you have to pay down on the price of the house
//RequiredPercentage - The percentage of property cost that needs to be a down payment
//Rent(decimal form) - The rent on the down payment if you have to take a loan to reach the required down payment.
//Amortization(decimal form) - The required Amortization percentage on the down payment you hade to take a loan for.
type DownPayment struct {
	AmountInHand       float64
	RequiredPercentage float64
	Rent               float64
	Amortization       float64
}

//RentRebate describes the rent rebate you get
//on loans. It's modelled after how the Swedish
//rent rebate system works.
//BeforeLimit(decimal form) - Rent rebate percentage before you reach the rent cost limit (over a year)
//AfterLimit(decimal form) - Rent rebate percentage after you reach the rent cost limit
//Limit - Limit where you payed enough rentcosts to go to the AfterLimit rent rebate percentage
type RentRebate struct {
	Limit       float64
	BeforeLimit float64
	AfterLimit  float64
}

//PropertyTax modelled after the Swedish system
//TaxationValuePercentageOfValue(decimal form) - Is the stupid percentage of the market value of your house which counts to the taxation value. Its complicated in how the market price is set, here we approximate it to be the price of the house.
//Percent(decimal form) -  is the tax you need to pay on the taxation value of the house
//Roof - the maximum value the property tax can reach.
type PropertyTax struct {
	TaxationValuePercentageOfValue float64
	Percent                        float64
	Roof                           float64
}

//Rent(decimal form)
//Amortization(decimal form)
//DownPayment - self explanatory
type Mortgage struct {
	Rent         float64
	Amortization float64
	DownPayment  DownPayment
}

func RequiredDownPayment(price float64, downPayment DownPayment) float64 {
	return price * downPayment.RequiredPercentage
}

//Hmm need to solve issue of rebateandtax being yearly if not jämkad.
func CalculateMonthly(price, operatingCostMonthly float64, mortgage Mortgage, rentRebate RentRebate, propertyTax PropertyTax, propertyInsuranceMonthly float64) (AmortizationMonthly, RealCostMonthly float64, err error) {
	downPayment := mortgage.DownPayment

	downPaymentBorrowed := math.Max(0, RequiredDownPayment(price, downPayment)-downPayment.AmountInHand)

	mortgageAmount := price - (mortgage.DownPayment.AmountInHand + downPaymentBorrowed)
	if mortgageAmount < 0 {
		return 0, 0, errors.New("Can not down pay more then the price of the property")
	}

	mainRent := mortgageAmount * mortgage.Rent
	downPaymentRent := downPaymentBorrowed * downPayment.Rent
	rent := mainRent + downPaymentRent

	rebate := math.Min(rent, rentRebate.Limit)*rentRebate.BeforeLimit +
		math.Max(0, rent-rentRebate.Limit)*rentRebate.AfterLimit

	rentCost := rent - rebate

	mainAmortization := (mortgageAmount) * mortgage.Amortization

	dpAmortizationCost := (downPaymentBorrowed) * downPayment.Amortization

	propertyTaxCost := math.Min(price*propertyTax.TaxationValuePercentageOfValue*propertyTax.Percent, propertyTax.Roof) / 12.0

	return rentCost/12 + operatingCostMonthly + propertyInsuranceMonthly + propertyTaxCost,
		mainAmortization/12 + dpAmortizationCost/12,
		nil
}

//Takes the price of the house with the mortgage deed currently on the house and relevant tax percentage
//for mortgage deeds and title deeds, and returns the total extra cost at purchase.
//
//taxes are given i a decimal percentage i.e 50% = 0.5
func ExtraAtPurchase(price, mortgageDeedCurrent, mortgageDeedTax, titleDeedTax float64) float64 {
	mortgageCost := (price - mortgageDeedCurrent) * mortgageDeedTax
	titleDeedCost := price * titleDeedTax
	return mortgageCost + titleDeedCost
}
