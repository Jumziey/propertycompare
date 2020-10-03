package propertydb

type PropertyType uint8

const (
	House = PropertyType(iota + 12) //12 just a random number
	Condo
)

func (pt *PropertyType) Valid() bool {
	return *pt == House || *pt == Condo
}

type City string
type StreetAddress string

type Info struct {
	PriceAsking              float64
	PriceFinal               float64
	Type                     PropertyType
	OperatingCosts           float64
	PropertyInsuranceMonthly float64
	CurrentMortgageDeed      float64
}

type Listing struct {
	City          City
	StreetAddress StreetAddress
	Info          Info
}
