package propertycost

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHouseMonthly(t *testing.T) {
	price := float64(1500000)
	operatingCostMonthly := float64(12000)
	propertyEnsuranceMonthly := float64(2000)
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
	taxProperty := TaxProperty{
		TaxationValuePercentageOfValue: 1,
		Percent:                        0.0075,
		Roof:                           8439,
	}

	expectedAmortizationMonthly := float64(7640)
	expectedRealCostMonthly := float64(8595.08333333)

	realCostMonthly, amortizationMonthly, _ := HouseMonthly(price, operatingCostMonthly, mortgage, rentRebate, taxProperty, propertyEnsuranceMonthly)
	assert.InDelta(t, expectedRealCostMonthly, realCostMonthly, 0.001, "RealCostMonthly")
	assert.InDelta(t, expectedAmortizationMonthly, amortizationMonthly, 0.0001, "Amortization")
}

func TestHouseMonthlyMoreDownPaymentThenRequired(t *testing.T) {
	price := float64(1500000)
	operatingCostMonthly := float64(12000)
	propertyEnsuranceMonthly := float64(2000)
	mortgage := Mortgage{
		Rent:         0.035,
		Amortization: 0.05,
		DownPayment: DownPayment{
			AmountInHand:       250000,
			RequiredPercentage: 0.16, //Should not matter
			Rent:               0.14, //Should not matter
			Amortization:       0.12, //Should not matter
		},
	}
	rentRebate := RentRebate{
		Limit:       100000,
		BeforeLimit: 0.3,
		AfterLimit:  0.21,
	}
	taxProperty := TaxProperty{
		TaxationValuePercentageOfValue: 1,
		Percent:                        0.0075,
		Roof:                           8439,
	}

	expectedRealCostMonthly := float64(6255.333333)
	expectedAmortizationMonthly := float64(5208.3333333)

	realCostMonthly, amortizationMonthly, _ := HouseMonthly(price, operatingCostMonthly, mortgage, rentRebate, taxProperty, propertyEnsuranceMonthly)
	assert.InDelta(t, expectedRealCostMonthly, realCostMonthly, 0.001, "RealCostMonthly")
	assert.InDelta(t, expectedAmortizationMonthly, amortizationMonthly, 0.0001, "Amortization")
}

func TestHouseMonthlyTooMuchDownPayment(t *testing.T) {
	price := float64(1500000)
	operatingCostMonthly := float64(12000)
	propertyEnsuranceMonthly := float64(2000)
	mortgage := Mortgage{
		Rent:         0.04,
		Amortization: 0.05,
		DownPayment: DownPayment{
			AmountInHand:       price + 1,
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
	taxProperty := TaxProperty{
		TaxationValuePercentageOfValue: 0.75,
		Percent:                        0.0075,
		Roof:                           8439,
	}

	_, _, err := HouseMonthly(price, operatingCostMonthly, mortgage, rentRebate, taxProperty, propertyEnsuranceMonthly)
	assert.NotNil(t, err)
}

func TestHouseMonthlyRentRebateAboveLimit(t *testing.T) {
	price := float64(1500000)
	operatingCostMonthly := float64(12000)
	propertyEnsuranceMonthly := float64(2000)
	mortgage := Mortgage{
		Rent:         0.10,
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
	taxProperty := TaxProperty{
		TaxationValuePercentageOfValue: 1,
		Percent:                        0.0075,
		Roof:                           8439,
	}

	expectedAmortizationMonthly := float64(7640)
	expectedRealCostMonthly := float64(13451.033333333333)

	realCostMonthly, amortizationMonthly, _ := HouseMonthly(price, operatingCostMonthly, mortgage, rentRebate, taxProperty, propertyEnsuranceMonthly)
	assert.InDelta(t, expectedRealCostMonthly, realCostMonthly, 0.001, "RealCostMonthly")
	assert.InDelta(t, expectedAmortizationMonthly, amortizationMonthly, 0.0001, "Amortization")
}

func TestHouseMonthlyTaxPropertyBelowRoof(t *testing.T) {
	price := float64(1000000)
	operatingCost := float64(12000)
	propertyEnsuranceMonthly := float64(2000)
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
	taxProperty := TaxProperty{
		TaxationValuePercentageOfValue: 0.75,
		Percent:                        0.0075,
		Roof:                           8439,
	}

	expectedRealCostMonthly := float64(6727.25)
	expectedAmortizationMonthly := float64(5090)

	realCostMonthly, amortizationMonthly, _ := HouseMonthly(price, operatingCost, mortgage, rentRebate, taxProperty, propertyEnsuranceMonthly)
	assert.InDelta(t, expectedRealCostMonthly, realCostMonthly, 0.001, "RealCostMonthly")
	assert.InDelta(t, expectedAmortizationMonthly, amortizationMonthly, 0.0001, "Amortization")
}

func TestHousePurchaseFees(t *testing.T) {
	price := float64(100000)
	mortgageDeedCurrent := float64(10000)
	mortgageDeedTax := float64(0.1)
	titleDeedTax := float64(0.2)
	expectedFees := float64(29000)

	totalExtraCost := HousePurchaseFees(price, mortgageDeedCurrent, mortgageDeedTax, titleDeedTax)

	assert.Equal(t, expectedFees, totalExtraCost)
}

func TestCondoMonthly(t *testing.T) {
	price := float64(1500000)
	operatingCostMonthly := float64(12000)
	propertyEnsuranceMonthly := float64(2000)
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

	expectedAmortizationMonthly := float64(7640)
	expectedRealCostMonthly := float64(7891.83333333)

	realCostMonthly, amortizationMonthly, _ := CondoMonthly(price, operatingCostMonthly, mortgage, rentRebate, propertyEnsuranceMonthly)
	assert.InDelta(t, expectedRealCostMonthly, realCostMonthly, 0.001, "RealCostMonthly")
	assert.InDelta(t, expectedAmortizationMonthly, amortizationMonthly, 0.0001, "Amortization")
}

func TestCondoMonthlyMoreDownPaymentThenRequired(t *testing.T) {
	price := float64(1500000)
	operatingCostMonthly := float64(12000)
	propertyEnsuranceMonthly := float64(2000)
	mortgage := Mortgage{
		Rent:         0.035,
		Amortization: 0.05,
		DownPayment: DownPayment{
			AmountInHand:       250000,
			RequiredPercentage: 0.16, //Should not matter
			Rent:               0.14, //Should not matter
			Amortization:       0.12, //Should not matter
		},
	}
	rentRebate := RentRebate{
		Limit:       100000,
		BeforeLimit: 0.3,
		AfterLimit:  0.21,
	}

	expectedRealCostMonthly := float64(5552.083333)
	expectedAmortizationMonthly := float64(5208.3333333)

	realCostMonthly, amortizationMonthly, _ := CondoMonthly(price, operatingCostMonthly, mortgage, rentRebate, propertyEnsuranceMonthly)
	assert.InDelta(t, expectedRealCostMonthly, realCostMonthly, 0.001, "RealCostMonthly")
	assert.InDelta(t, expectedAmortizationMonthly, amortizationMonthly, 0.0001, "Amortization")
}

func TestCondoMonthlyTooMuchDownPayment(t *testing.T) {
	price := float64(1500000)
	operatingCostMonthly := float64(12000)
	propertyEnsuranceMonthly := float64(2000)
	mortgage := Mortgage{
		Rent:         0.04,
		Amortization: 0.05,
		DownPayment: DownPayment{
			AmountInHand:       price + 1,
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

	_, _, err := CondoMonthly(price, operatingCostMonthly, mortgage, rentRebate, propertyEnsuranceMonthly)
	assert.NotNil(t, err)
}

func TestCondoMonthlyRentRebateAboveLimit(t *testing.T) {
	price := float64(1500000)
	operatingCostMonthly := float64(12000)
	propertyEnsuranceMonthly := float64(2000)
	mortgage := Mortgage{
		Rent:         0.10,
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

	expectedAmortizationMonthly := float64(7640)
	expectedRealCostMonthly := float64(12747.783333333333)

	realCostMonthly, amortizationMonthly, _ := CondoMonthly(price, operatingCostMonthly, mortgage, rentRebate, propertyEnsuranceMonthly)
	assert.InDelta(t, expectedRealCostMonthly, realCostMonthly, 0.001, "RealCostMonthly")
	assert.InDelta(t, expectedAmortizationMonthly, amortizationMonthly, 0.0001, "Amortization")
}
