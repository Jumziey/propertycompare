package propertycost

import (
	"errors"
	"math"
)

//### NOTES ####
//TaxProperty (actual value)
//Calculated as
//TaxPropertyValue = 75 percent of price
//(not a true value, normally its decided on the average costs
// of properties in your neighbourhood over the last 3 years.)
//TaxProperty = min(8349, TaxPropertyValue*0.0075)
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

//TaxProperty modelled after the Swedish system
//TaxationValuePercentageOfValue(decimal form) - Is the stupid percentage of the market value of your house which counts to the taxation value. Its complicated in how the market price is set, here we approximate it to be the price of the house.
//Percent(decimal form) -  is the tax you need to pay on the taxation value of the house
//Roof - the maximum value the property tax can reach.
type TaxProperty struct {
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

//Hmm need to solve issue of rebateandtax being yearly if not jämkad.
func CalculateMonthly(price, operatingCostMonthly float64, mortgage Mortgage, rentRebate RentRebate, taxProperty TaxProperty, propertyInsuranceMonthly float64) (AmortizationMonthly, RealCostMonthly float64, err error) {

	mainRent, downPaymentRent, err := Rent(price, mortgage)
	if err != nil {
		return 0, 0, err
	}
	rent := mainRent + downPaymentRent
	rentCost := rent - Rebate(rent, rentRebate)

	mainAmortization, dpAmortization, err := Amortization(price, mortgage)
	if err != nil {
		return 0, 0, err
	}

	return rentCost/12 + operatingCostMonthly + propertyInsuranceMonthly + TaxPropertyCost(price, taxProperty)/12.0,
		mainAmortization/12 + dpAmortization/12,
		nil
}

func TaxPropertyCost(price float64, taxProperty TaxProperty) float64 {
	return math.Min(price*taxProperty.TaxationValuePercentageOfValue*taxProperty.Percent, taxProperty.Roof)
}

func Rebate(rentTotal float64, rentRebate RentRebate) float64 {
	return math.Min(rentTotal, rentRebate.Limit)*rentRebate.BeforeLimit +
		math.Max(0, rentTotal-rentRebate.Limit)*rentRebate.AfterLimit
}

func Amortization(price float64, mortgage Mortgage) (mainAmortization float64, downPaymentAmortization float64, err error) {

	mortgageTot, err := mortgageTotal(price, mortgage)
	if err != nil {
		return 0, 0, err
	}

	mainAmortization = (mortgageTot) * mortgage.Amortization
	downPaymentAmortization = downPaymentBorrowed(price, mortgage.DownPayment) * mortgage.DownPayment.Amortization

	return mainAmortization, downPaymentAmortization, nil
}

func RequiredDownPayment(price float64, downPayment DownPayment) float64 {
	return price * downPayment.RequiredPercentage
}

func downPaymentBorrowed(price float64, downPayment DownPayment) float64 {
	return math.Max(0, RequiredDownPayment(price, downPayment)-downPayment.AmountInHand)
}

func Rent(price float64, mortgage Mortgage) (mainRent float64, downPaymentRent float64, err error) {
	downPayment := mortgage.DownPayment

	mortgageTot, err := mortgageTotal(price, mortgage)
	if err != nil {
		return 0, 0, err
	}

	mainRent = mortgageTot * mortgage.Rent
	downPaymentRent = downPaymentBorrowed(price, mortgage.DownPayment) * downPayment.Rent
	return mainRent, downPaymentRent, nil
}

func mortgageTotal(price float64, mortgage Mortgage) (float64, error) {
	mortgageAmount := price - (mortgage.DownPayment.AmountInHand + downPaymentBorrowed(price, mortgage.DownPayment))
	if mortgageAmount < 0 {
		return 0, errors.New("Can not down pay more then the price of the property")
	}
	return mortgageAmount, nil
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
