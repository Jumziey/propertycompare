package propertycost

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
	AmountInHand       float32
	RequiredPercentage float32
	Rent               float32
	Amortization       float32
}

//RentRebate describes the rent rebate you get
//on loans. It's modelled after how the Swedish
//rent rebate system works.
//BeforeLimit(decimal form) - Rent rebate percentage before you reach the rent cost limit (over a year)
//AfterLimit(decimal form) - Rent rebate percentage after you reach the rent cost limit
//Limit - Limit where you payed enough rentcosts to go to the AfterLimit rent rebate percentage
type RentRebate struct {
	Limit       float32
	BeforeLimit float32
	AfterLimit  float32
}

//PropertyTax modelled after the Swedish system
//TaxationValuePercentageOfValue(decimal form) - Is the stupid percentage of the market value of your house which counts to the taxation value. Its complicated in how the market price is set, here we approximate it to be the price of the house.
//Percent(decimal form) -  is the tax you need to pay on the taxation value of the house
//Roof - the maximum value the property tax can reach.
type PropertyTax struct {
	TaxationValuePercentageOfValue float32
	Percent                        float32
	Roof                           float32
}

//Rent(decimal form)
//Amortization(decimal form)
//DownPayment - self explanatory
type Mortgage struct {
	Rent         float32
	Amortization float32
	DownPayment  DownPayment
}

func CalculateMonthly(price, operatingCostMonthly float32, mortgage Mortgage, rentRebate RentRebate, propertyTax PropertyTax, propertyEnsuranceMonthly float32) (AmortizationMonthly, RealCostMonthly float32, err error) {
	downPayment := mortgage.DownPayment

	requiredDownPayment := price * downPayment.RequiredPercentage

	mainRentCost := (price - requiredDownPayment) * mortgage.Rent * (1.0 - rentRebate.BeforeLimit)
	mainAmortization := (price - requiredDownPayment) * mortgage.Amortization

	dpRentCost := (requiredDownPayment - downPayment.AmountInHand) * downPayment.Rent * (1.0 - rentRebate.BeforeLimit)
	dpAmortizationCost := (requiredDownPayment - downPayment.AmountInHand) * downPayment.Amortization

	propertyTaxCost := propertyTax.Roof / 12.0

	return mainRentCost/12 + dpRentCost/12 + operatingCostMonthly + propertyEnsuranceMonthly + propertyTaxCost,
		mainAmortization/12 + dpAmortizationCost/12,
		nil
}

//Takes the price of the house with the mortgage deed currently on the house and relevant tax percentage
//for mortgage deeds and title deeds, and returns the total extra cost at purchase.
//
//taxes are given i a decimal percentage i.e 50% = 0.5
func ExtraAtPurchase(price, mortgageDeedCurrent, mortgageDeedTax, titleDeedTax float32) float32 {
	mortgageCost := (price - mortgageDeedCurrent) * mortgageDeedTax
	titleDeedCost := price * titleDeedTax
	return mortgageCost + titleDeedCost
}
