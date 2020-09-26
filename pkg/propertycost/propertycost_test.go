package propertycost

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMonthly(t *testing.T) {
	price := float32(1500000)
	operatingCostMonthly := float32(1000)
	propertyEnsuranceMonthly := float32(2000)
	mortgage := Mortgage{
		Rent:         0.04,
		Amortization: 0.05,
		DownPayment: DownPayment{
			AmountInHand:       1000,
			RequiredPercentage: 0.16,
			Rent:               0.14,
			Amortization:       0.12,
		},
	}
	rentRebate := RentRebate{
		Limit:       100000,
		BeforeLimit: 0.3,
		AfterLimit:  0.21,
	}
	propertyTax := PropertyTax{
		Percent: 0.1,
		Roof:    8439,
	}

	expectedAmortizationMonthly := float32(7640)
	expectedRealCostMonthly := float32(8595.08333333)

	realCostMonthly, amortizationMontly, _ := CalculateMonthly(price, operatingCostMonthly, mortgage, rentRebate, propertyTax, propertyEnsuranceMonthly)
	assert.InDelta(t, expectedRealCostMonthly, realCostMonthly, 0.001, "RealCostMonthly")
	assert.InDelta(t, expectedAmortizationMonthly, amortizationMontly, 0.0001, "Amortization")
}

//Calculate if downPayment.AmountInHand is larger then requiredDownPayment

//Calculate if rent is more then 100 000k a year (so tax rebate on rent payments gets yanky 21% rule) 30% for the first 100kkr

//Calculate if price is less then 1.4 mkr (so propertytax gets yanky)

func TestExtraAtPurchase(t *testing.T) {
	price := float32(100000)
	mortgageDeedCurrent := float32(10000)
	mortgageDeedTax := float32(0.1)
	titleDeedTax := float32(0.2)
	expectedTotalExtraCost := float32(29000)

	totalExtraCost := ExtraAtPurchase(price, mortgageDeedCurrent, mortgageDeedTax, titleDeedTax)

	assert.Equal(t, expectedTotalExtraCost, totalExtraCost)
}
