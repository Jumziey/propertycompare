package propertydb

import "testing"

func ValidateAddAndShow(t *testing.T) {
	//Add a valid property, make sure no error
	//Check city on Show
}
func ValidateAddSameTwiceError(t *testing.T) {
	//Add valid property twice, make sure an error Shows on second add
}

func ValidateCityAndStreetAddressIdentifies(t *testing.T) {
	//Add two valid properties in the same city but different addresses
	//make sure no error
	//Add two valid properties on the same address but different cities
	//make sure no error
}

func ValidateUpdate(t *testing.T) {
	//Add property, update property, validate update took place with Show
}

func ValidateList(t *testing.T) {
	//Add three properties, validate same with List()
}

func ValidateDelete(t *testing.T) {
	//Add three properties, delete one validate with List()
}
