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
	expectedListing := defaultListing()

	err := pdb.Add(expectedListing.City, expectedListing.StreetAddress, expectedListing.Info)
	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
	err = pdb.Add(expectedListing.City, expectedListing.StreetAddress, expectedListing.Info)
	assert.NotNil(t, err, ".Add(...) does not return error when trying to add an existing property")
}

func ValidateCityAndStreetAddressIdentifies(t *testing.T, pdb PropertyDB) {
	listingFirst := defaultListing()
	listingSecond := Listing{
		City:          City("SecretTown"),
		StreetAddress: StreetAddress("Secret Address"),
		Info: Info{
			PriceAsking:              1212121,
			PriceFinal:               14141414,
			Type:                     House,
			OperatingCosts:           1211,
			PropertyInsuranceMonthly: 121,
			CurrentMortgageDeed:      1238,
			Notes:                    "Hello mr sunshine",
		},
	}
	err := pdb.Add(listingFirst.City, listingFirst.StreetAddress, listingFirst.Info)
	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
	err = pdb.Add(listingSecond.City, listingSecond.StreetAddress, listingSecond.Info)
	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")

	retSecond, err := pdb.Show(listingSecond.City, listingSecond.StreetAddress)
	assert.Equal(t, listingSecond, retSecond)
	retFirst, err := pdb.Show(listingFirst.City, listingFirst.StreetAddress)
	assert.Equal(t, listingFirst, retFirst)
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
