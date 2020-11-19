package property

//City is the name of a city
type City string

//StreetAddress is a street address with name and number
type StreetAddress string

//Type is a type of property
type Type uint8

const (
	//Unknown is an uninitialized value
	Unknown = Type(iota)
	//House is a self owned house.
	House
	//Condo is an owned rental
	Condo
)

//Valid just returns a bool describing if the
//type value is valid.
func (pt *Type) Valid() bool {
	return *pt == House || *pt == Condo
}

//Property contains all relevant information
//of a property that's used to calculate
//everything else and index the property
type Property struct {
	//City is the name of the city the property exist in
	City City
	//StreetAddress is the street address of the property
	StreetAddress StreetAddress
	//PriceAsking was the asking price for the property
	PriceAsking float64
	//PriceFinal the final price of the property
	PriceFinal float64
	//Type is the type of property
	Type Type
	//OperatingCosts is the operating cost of the property yearly
	OperatingCosts float64 //hmm
	//PropertyInsuranceMonthly is an insurance quote on the property
	PropertyInsuranceMonthly float64 //hmm
	//MortgageAmount is the current mortgage size on the property
	MortgageAmount float64
	//Notes is just general notes aboat the property
	Notes string
}
