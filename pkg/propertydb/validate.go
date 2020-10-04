package propertydb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func defaultCity() City {
	return City("Town")
}

func defaultAddress() StreetAddress {
	return StreetAddress("Address 1")
}

func defaultInfo() Info {
	return Info{
		PriceAsking:              12,
		PriceFinal:               14,
		Type:                     House,
		OperatingCosts:           1231,
		PropertyInsuranceMonthly: 121,
		CurrentMortgageDeed:      1238,
		Notes:                    "Hello mr sunshine",
	}
}

func defaultListing() Listing {
	return Listing{
		City:          defaultCity(),
		StreetAddress: defaultAddress(),
		Info:          defaultInfo(),
	}
}

func ValidateAddAndShow(t *testing.T, pdb PropertyDB) {
	expectedListing := defaultListing()

	err := pdb.Add(expectedListing.City, expectedListing.StreetAddress, expectedListing.Info)
	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")

	listing, err := pdb.Show(expectedListing.City, expectedListing.StreetAddress)
	assert.Nil(t, err, ".Show(...) returns error when trying to add a property")

	assert.Equal(t, expectedListing, listing)
}

func ValidateAddSameTwiceError(t *testing.T, pdb PropertyDB) {
	//Add valid property twice, make sure an error Shows on second add
}

func ValidateCityAndStreetAddressIdentifies(t *testing.T, pdb PropertyDB) {
	//Add two valid properties in the same city but different addresses
	//make sure no error
	//Add two valid properties on the same address but different cities
	//make sure no error
}

func ValidateUpdate(t *testing.T, pdb PropertyDB) {
	//Add property, update property, validate update took place with Show
}

func ValidateList(t *testing.T, pdb PropertyDB) {
	//Add three properties, validate same with List()
}

func ValidateDelete(t *testing.T, pdb PropertyDB) {
	//Add three properties, delete one validate with List()
}
