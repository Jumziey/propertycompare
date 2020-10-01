package propertycost

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMonthly(t *testing.T) {
	price := float64(1500000)
	operatingCostMonthly := float64(1000)
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

func TestCalculateMonthlyMoreDownPaymentThenRequired(t *testing.T) {
	price := float64(1500000)
	operatingCostMonthly := float64(1000)
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

func TestCalculateMonthlyTooMuchDownPayment(t *testing.T) {
	price := float64(1500000)
	operatingCostMonthly := float64(1000)
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

func TestCalculateMonthlyRentRebateAboveLimit(t *testing.T) {
	price := float64(1500000)
	operatingCostMonthly := float64(1000)
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

func TestCalculateMonthlyTaxPropertyBelowRoof(t *testing.T) {
	price := float64(1000000)
	operatingCostMonthly := float64(1000)
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

	realCostMonthly, amortizationMonthly, _ := HouseMonthly(price, operatingCostMonthly, mortgage, rentRebate, taxProperty, propertyEnsuranceMonthly)
	assert.InDelta(t, expectedRealCostMonthly, realCostMonthly, 0.001, "RealCostMonthly")
	assert.InDelta(t, expectedAmortizationMonthly, amortizationMonthly, 0.0001, "Amortization")
}

func TestHouseExtraAtPurchase(t *testing.T) {
	price := float64(100000)
	mortgageDeedCurrent := float64(10000)
	mortgageDeedTax := float64(0.1)
	titleDeedTax := float64(0.2)
	expectedTotalExtraCost := float64(29000)

	totalExtraCost := HousePurchaseFees(price, mortgageDeedCurrent, mortgageDeedTax, titleDeedTax)

	assert.Equal(t, expectedTotalExtraCost, totalExtraCost)
}
