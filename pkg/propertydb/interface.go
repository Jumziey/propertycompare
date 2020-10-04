package propertydb

type PropertyDB interface {
	Add(City, StreetAddress, Info) error //more
	Delete(City, StreetAddress) error
	Update(City, StreetAddress, Info) error
	List() ([]Listing, error) //more
	Show(City, StreetAddress) (Listing, error)
}
