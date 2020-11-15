// package property
//
// import (
// 	"testing"
//
// 	"github.com/stretchr/testify/assert"
// )
//
// func defaultCity() City {
// 	return City("Town")
// }
//
// func defaultAddress() StreetAddress {
// 	return StreetAddress("Address 1")
// }
//
// func defaultInfo() Info {
// 	return Info{
// 		PriceAsking:              12,
// 		PriceFinal:               14,
// 		Type:                     House,
// 		OperatingCosts:           1231,
// 		PropertyInsuranceMonthly: 121,
// 		CurrentMortgageDeed:      1238,
// 		Notes:                    "Hello mr sunshine",
// 	}
// }
//
// func defaultListing() Listing {
// 	return Listing{
// 		City:          defaultCity(),
// 		StreetAddress: defaultAddress(),
// 		Info:          defaultInfo(),
// 	}
// }
//
// func ValidateAddAndShow(t *testing.T, pdb PropertyDB) {
// 	expectedListing := defaultListing()
//
// 	err := pdb.Add(expectedListing.City, expectedListing.StreetAddress, expectedListing.Info)
// 	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
//
// 	listing, err := pdb.Show(expectedListing.City, expectedListing.StreetAddress)
// 	assert.Nil(t, err, ".Show(...) returns error when trying to show an added property")
//
// 	assert.Equal(t, expectedListing, listing)
// }
//
// func ValidateAddSameTwiceError(t *testing.T, pdb PropertyDB) {
// 	expectedListing := defaultListing()
//
// 	err := pdb.Add(expectedListing.City, expectedListing.StreetAddress, expectedListing.Info)
// 	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
// 	err = pdb.Add(expectedListing.City, expectedListing.StreetAddress, expectedListing.Info)
// 	assert.NotNil(t, err, ".Add(...) does not return error when trying to add an existing property")
// }
//
// func ValidateCityAndStreetAddressIdentifies(t *testing.T, pdb PropertyDB) {
// 	listingFirst := defaultListing()
// 	listingSecond := Listing{
// 		City:          City("SecretTown"),
// 		StreetAddress: StreetAddress("Secret Address"),
// 		Info: Info{
// 			PriceAsking:              1212121,
// 			PriceFinal:               14141414,
// 			Type:                     House,
// 			OperatingCosts:           1211,
// 			PropertyInsuranceMonthly: 121,
// 			CurrentMortgageDeed:      1238,
// 			Notes:                    "Hello mr sunshine",
// 		},
// 	}
// 	err := pdb.Add(listingFirst.City, listingFirst.StreetAddress, listingFirst.Info)
// 	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
// 	err = pdb.Add(listingSecond.City, listingSecond.StreetAddress, listingSecond.Info)
// 	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
//
// 	retSecond, err := pdb.Show(listingSecond.City, listingSecond.StreetAddress)
// 	assert.Nil(t, err, ".Show(...) returns error when trying to show the second added property")
// 	assert.Equal(t, listingSecond, retSecond)
// 	retFirst, err := pdb.Show(listingFirst.City, listingFirst.StreetAddress)
// 	assert.Nil(t, err, ".Show(...) returns error when trying to show the first added property")
// 	assert.Equal(t, listingFirst, retFirst)
// }
//
// func ValidateUpdate(t *testing.T, pdb PropertyDB) {
// 	//Add property, update property, validate update took place with Show
//
// 	listing := defaultListing()
// 	err := pdb.Add(listing.City, listing.StreetAddress, listing.Info)
// 	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
//
// 	updatedListing := listing
// 	updatedListing.Info.Notes = "I update and update"
//
// 	pdb.Update(updatedListing.City, updatedListing.StreetAddress, updatedListing.Info)
//
// 	l, err := pdb.Show(listing.City, listing.StreetAddress)
// 	assert.Nil(t, err, ".Show(...) returns error when trying to read the updated listing")
// 	assert.Equal(t, updatedListing, l)
// }
//
// func ValidateList(t *testing.T, pdb PropertyDB) {
// 	//Add three properties, validate same with List()
// 	l1 := defaultListing()
// 	l2 := Listing{
// 		City:          City("SecretTown"),
// 		StreetAddress: StreetAddress("Secret Address"),
// 		Info: Info{
// 			PriceAsking:              1212121,
// 			PriceFinal:               14141414,
// 			Type:                     House,
// 			OperatingCosts:           1211,
// 			PropertyInsuranceMonthly: 121,
// 			CurrentMortgageDeed:      1238,
// 			Notes:                    "Hello mr sunshine",
// 		},
// 	}
// 	l3 := Listing{
// 		City:          City("ChristmasTown"),
// 		StreetAddress: StreetAddress("Christmas Address"),
// 		Info: Info{
// 			PriceAsking:              1212111,
// 			PriceFinal:               141914,
// 			Type:                     Condo,
// 			OperatingCosts:           1211,
// 			PropertyInsuranceMonthly: 131,
// 			CurrentMortgageDeed:      1228,
// 			Notes:                    "Snow and Snow",
// 		},
// 	}
//
// 	err := pdb.Add(l1.City, l1.StreetAddress, l1.Info)
// 	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
// 	err = pdb.Add(l2.City, l2.StreetAddress, l2.Info)
// 	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
// 	err = pdb.Add(l3.City, l3.StreetAddress, l3.Info)
// 	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
//
// 	listings, err := pdb.List()
// 	assert.Nil(t, err, ".List(...) returns error when trying to list properties")
// 	for _, l := range listings {
// 		if l1 != l && l2 != l && l3 != l {
// 			t.Log("l1", l1)
// 			t.Log("l2", l2)
// 			t.Log("l3", l3)
// 			t.Log("l", l)
// 			t.Log(".List() return faulty listings")
// 			t.Fail()
// 		}
//
// 	}
// }
//
// func ValidateDelete(t *testing.T, pdb PropertyDB) {
// 	//Add three properties, delete one, validate with List()
// 	l1 := defaultListing()
// 	l2 := Listing{
// 		City:          City("SecretTown"),
// 		StreetAddress: StreetAddress("Secret Address"),
// 		Info: Info{
// 			PriceAsking:              1212121,
// 			PriceFinal:               14141414,
// 			Type:                     House,
// 			OperatingCosts:           1211,
// 			PropertyInsuranceMonthly: 121,
// 			CurrentMortgageDeed:      1238,
// 			Notes:                    "Hello mr sunshine",
// 		},
// 	}
// 	l3 := Listing{
// 		City:          City("ChristmasTown"),
// 		StreetAddress: StreetAddress("Christmas Address"),
// 		Info: Info{
// 			PriceAsking:              1212111,
// 			PriceFinal:               141914,
// 			Type:                     Condo,
// 			OperatingCosts:           1211,
// 			PropertyInsuranceMonthly: 131,
// 			CurrentMortgageDeed:      1228,
// 			Notes:                    "Snow and Snow",
// 		},
// 	}
//
// 	err := pdb.Add(l1.City, l1.StreetAddress, l1.Info)
// 	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
// 	err = pdb.Add(l2.City, l2.StreetAddress, l2.Info)
// 	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
// 	err = pdb.Add(l3.City, l3.StreetAddress, l3.Info)
// 	assert.Nil(t, err, ".Add(...) returns error when trying to add a property")
//
// 	err = pdb.Delete(l2.City, l2.StreetAddress)
// 	assert.Nil(t, err, ".Delete(...) returns error when trying to delete a property")
//
// 	listings, err := pdb.List()
// 	assert.Nil(t, err, ".List(...) returns error when trying to list properties")
// 	for _, l := range listings {
// 		if l1 != l && l3 != l {
// 			t.Log("l1", l1)
// 			t.Log("l2", l2)
// 			t.Log("l3", l3)
// 			t.Log("l", l)
// 			t.Log(".Delete() seem to interact weirdly with .List()")
//
// 		}
// 	}
// }
