package property

//Store describes an interface to a type which handles
//a collection of properties.
type Store interface {
	Add(Property) error
	Delete(City, StreetAddress) error
	//Update the existing property with the Property.City and Property.StreetAddress
	//of the given Property.
	Update(Property) error
	//Returns a slice with all the properties in the Store.
	List() ([]Property, error) //more
	//Returns a specific property on the City and StreetAddress given
	Show(City, StreetAddress) (Property, error)
}
