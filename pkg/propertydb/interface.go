package propertydb

type PropertyDB interface {
	Add(City, StreetAddress, Info) (Info, error) //more
	Delete(City, StreetAddress) (Info, error)
	Update(City, StreetAddress, Info) (Info, error)
	List() ([]Listing, error) //more
	Show(City, StreetAddress) (Info, error)
}
