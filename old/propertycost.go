package propertycost

import (
	"errors"
	"math"
)

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

func HouseMonthly(price, operatingCost float64, mortgage Mortgage, rentRebate RentRebate, taxProperty TaxProperty, propertyInsuranceMonthly float64) (RealCostMonthly, AmortizationMonthly float64, err error) {
	condoCostMonthly, amortization, err := CondoMonthly(price, operatingCost, mortgage, rentRebate, propertyInsuranceMonthly)
	if err != nil {
		return 0, 0, err
	}
	return condoCostMonthly + HouseTax(price, taxProperty)/12.0,
		amortization,
		nil
}

func HouseTax(price float64, taxProperty TaxProperty) float64 {
	return math.Min(price*taxProperty.TaxationValuePercentageOfValue*taxProperty.Percent, taxProperty.Roof)
}

func CondoMonthly(price, operatingCost float64, mortgage Mortgage, rentRebate RentRebate, propertyInsuranceMonthly float64) (RealCostMonthly, AmortizationMonthly float64, err error) {

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

	return rentCost/12 + operatingCost/float64(12) + propertyInsuranceMonthly,
		mainAmortization/12 + dpAmortization/12,
		nil
}

func RequiredDownPayment(price float64, downPayment DownPayment) float64 {
	return price * downPayment.RequiredPercentage
}

//Yearly
func Rebate(rentTotal float64, rentRebate RentRebate) float64 {
	return math.Min(rentTotal, rentRebate.Limit)*rentRebate.BeforeLimit +
		math.Max(0, rentTotal-rentRebate.Limit)*rentRebate.AfterLimit
}

//Yearly
func Amortization(price float64, mortgage Mortgage) (mainAmortization float64, downPaymentAmortization float64, err error) {

	mortgageTot, err := mortgageTotal(price, mortgage)
	if err != nil {
		return 0, 0, err
	}

	mainAmortization = (mortgageTot) * mortgage.Amortization
	downPaymentAmortization = downPaymentBorrowed(price, mortgage.DownPayment) * mortgage.DownPayment.Amortization

	return mainAmortization, downPaymentAmortization, nil
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

//taxes are given i a decimal percentage i.e 50% = 0.5
func HousePurchaseFees(price, mortgageDeedCurrent, mortgageDeedTax, titleDeedTax float64) float64 {
	return MortgageDeed(price, mortgageDeedCurrent, mortgageDeedTax) + TitleDeed(price, titleDeedTax)
}

//taxes are given i a decimal percentage i.e 50% = 0.5
func MortgageDeed(price, mortgageDeedCurrent, mortgageDeedTax float64) float64 {
	return (price - mortgageDeedCurrent) * mortgageDeedTax
}

//taxes are given i a decimal percentage i.e 50% = 0.5
func TitleDeed(price, titleDeedTax float64) float64 {
	return price * titleDeedTax
}

func downPaymentBorrowed(price float64, downPayment DownPayment) float64 {
	return math.Max(0, RequiredDownPayment(price, downPayment)-downPayment.AmountInHand)
}

func mortgageTotal(price float64, mortgage Mortgage) (float64, error) {
	mortgageAmount := price - (mortgage.DownPayment.AmountInHand + downPaymentBorrowed(price, mortgage.DownPayment))
	if mortgageAmount < 0 {
		return 0, errors.New("Can not down pay more then the price of the property")
	}
	return mortgageAmount, nil
}
